// Copyright 2013 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package juju

import (
	"fmt"
	"time"

	"launchpad.net/loggo"

	"launchpad.net/juju-core/environs"
	"launchpad.net/juju-core/environs/config"
	"launchpad.net/juju-core/environs/configstore"
	"launchpad.net/juju-core/errors"
	"launchpad.net/juju-core/state/api"
)

var logger = loggo.GetLogger("juju")

// The following are variables so that they can be
// changed by tests.
var (
	apiOpen            = api.Open
	apiClose           = (*api.State).Close
	defaultConfigStore = func() (environs.ConfigStorage, error) {
		return configstore.NewDisk(config.JujuHomePath("environments"))
	}
	providerConnectDelay = 2 * time.Second
)

// APIConn holds a connection to a juju environment and its
// associated state through its API interface.
type APIConn struct {
	Environ environs.Environ
	State   *api.State
}

// NewAPIConn returns a new Conn that uses the
// given environment. The environment must have already
// been bootstrapped.
func NewAPIConn(environ environs.Environ, dialOpts api.DialOpts) (*APIConn, error) {
	_, info, err := environ.StateInfo()
	if err != nil {
		return nil, err
	}
	info.Tag = "user-admin"
	password := environ.Config().AdminSecret()
	if password == "" {
		return nil, fmt.Errorf("cannot connect without admin-secret")
	}
	info.Password = password

	st, err := apiOpen(info, dialOpts)
	// TODO(rog): handle errUnauthorized when the API handles passwords.
	if err != nil {
		return nil, err
	}
	// TODO(rog): implement updateSecrets (see Conn.updateSecrets)
	return &APIConn{
		Environ: environ,
		State:   st,
	}, nil
}

// Close terminates the connection to the environment and releases
// any associated resources.
func (c *APIConn) Close() error {
	return c.State.Close()
}

// newAPIFromName implements the bulk of NewAPIClientFromName
// but is separate for testing purposes.
func newAPIFromName(envName string) (*api.State, error) {
	store, err := defaultConfigStore()
	if err != nil {
		return nil, err
	}
	// Try to read the default environment configuration file.
	// If it doesn't exist, we carry on in case
	// there's some environment info for that environment.
	// This enables people to copy environment files
	// into their .juju/environments directory and have
	// them be directly useful with no further configuration changes.
	envs, err := environs.ReadEnvirons("")
	if err == nil {
		if envName == "" {
			envName = envs.Default
		}
		if envName == "" {
			return nil, fmt.Errorf("no default environment found")
		}
	} else if !environs.IsNoEnv(err) {
		return nil, err
	}

	// Try to connect to the API concurrently using two
	// different possible sources of truth for the API endpoint.
	// Our preference is for the API endpoint cached in the
	// API info, because we know that without needing to
	// access any remote provider. However, the addresses
	// stored there may no longer be current (and the network
	// connection may take a very long time to time out)
	// so we also try to connect using information found
	// from the provider config. If we have some local
	// information, we only start to make that connection
	// after some suitable delay, so that in the hopefully
	// usual case, we will make the connection to the API
	// and never hit the provider.

	stop := make(chan struct{})
	defer close(stop)

	infoResult := apiInfoConnect(store, envName, stop)

	var cfgResult <-chan apiOpenResult
	if envs != nil {
		delay := providerConnectDelay
		if infoResult == nil {
			delay = 0
		}
		cfgResult = apiConfigConnect(envs, envName, stop, delay)
	}

	if infoResult == nil && cfgResult == nil {
		return nil, errors.NotFoundf("environment %q", envName)
	}
	var (
		st      *api.State
		infoErr error
		cfgErr  error
	)
	for st == nil && (infoResult != nil || cfgResult != nil) {
		select {
		case r := <-infoResult:
			st = r.st
			infoErr = r.err
			infoResult = nil
		case r := <-cfgResult:
			st = r.st
			cfgErr = r.err
			cfgResult = nil
		}
	}
	if st != nil {
		// One potential issue: there may still be a lingering
		// API connection, which will use resources until it
		// finally succeeds or fails. Unless we are making hundreds
		// of API connections, this is unlikely to be a problem.
		return st, nil
	}
	if cfgErr != nil {
		// Return the error from the configuration lookup if we
		// have one, because that information should be most current.
		logger.Warningf("discarding API open error: %v", infoErr)
		return nil, cfgErr
	}
	return nil, infoErr
}

func cleanUpAPIOpenResult(resultc <-chan apiOpenResult) {
	if resultc == nil {
		return
	}
	go func() {
		r := <-resultc
		if r.err != nil {
			logger.Warningf("disarding stale API open error: %v", r.err)
		} else {
			r.st.Close()
		}
	}()
}

// NewAPIClientFromName returns an api.Client connected to the API Server for
// the named environment. If envName is "", the default environment
// will be used.
func NewAPIClientFromName(envName string) (*api.Client, error) {
	st, err := newAPIFromName(envName)
	if err != nil {
		return nil, err
	}
	return st.Client(), nil
}

type apiOpenResult struct {
	st  *api.State
	err error
}

// apiInfoConnect looks for endpoint on the given environment and
// tries to connect to it, sending the result on the returned channel.
func apiInfoConnect(store environs.ConfigStorage, envName string, stop <-chan struct{}) <-chan apiOpenResult {
	resultc := make(chan apiOpenResult)
	info, err := store.ReadInfo(envName)
	if err != nil {
		if errors.IsNotFoundError(err) {
			return nil
		}
		go sendAPIOpenResult(resultc, stop, nil, err)
		return resultc
	}
	endpoint := info.APIEndpoint()
	if info == nil || len(endpoint.Addresses) == 0 {
		return nil
	}
	go func() {
		st, err := apiOpen(&api.Info{
			Addrs:    endpoint.Addresses,
			CACert:   []byte(endpoint.CACert),
			Tag:      "user-" + info.APICredentials().User,
			Password: info.APICredentials().Password,
		}, api.DefaultDialOpts())
		sendAPIOpenResult(resultc, stop, st, err)
	}()
	return resultc
}

func sendAPIOpenResult(resultc chan apiOpenResult, stop <-chan struct{}, st *api.State, err error) {
	select {
	case <-stop:
		if err != nil {
			logger.Warningf("disarding stale API open error: %v", err)
		} else {
			apiClose(st)
		}
	case resultc <- apiOpenResult{st, err}:
	}
}

// apiConfigConnect looks for configuration info on the given environment,
// and tries to use an Environ constructed from that to connect to
// its endpoint. It only starts the attempt after the given delay,
// to allow the faster apiInfoConnect to hopefully succeed first.
func apiConfigConnect(envs *environs.Environs, envName string, stop <-chan struct{}, delay time.Duration) <-chan apiOpenResult {
	cfg, err := envs.Config(envName)
	if errors.IsNotFoundError(err) {
		return nil
	}
	resultc := make(chan apiOpenResult)
	connect := func() (*api.State, error) {
		if err != nil {
			return nil, err
		}
		select {
		case <-time.After(delay):
		case <-stop:
			return nil, fmt.Errorf("aborted")
		}
		environ, err := environs.New(cfg)
		if err != nil {
			return nil, err
		}
		apiConn, err := NewAPIConn(environ, api.DefaultDialOpts())
		if err != nil {
			return nil, err
		}
		return apiConn.State, nil
	}
	go func() {
		st, err := connect()
		sendAPIOpenResult(resultc, stop, st, err)
	}()
	return resultc
}
