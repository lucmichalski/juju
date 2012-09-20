package state_test

import (
	. "launchpad.net/gocheck"
	"launchpad.net/juju-core/state"
)

type AssignSuite struct {
	ConnSuite
	charm   *state.Charm
	service *state.Service
	unit    *state.Unit
}

var _ = Suite(&AssignSuite{})

func (s *AssignSuite) SetUpTest(c *C) {
	s.ConnSuite.SetUpTest(c)
	s.charm = s.AddTestingCharm(c, "dummy")
	var err error
	s.service, err = s.State.AddService("wordpress", s.charm)
	c.Assert(err, IsNil)
	s.unit, err = s.service.AddUnit()
	c.Assert(err, IsNil)
	// Create root machine that shouldn't be used unless requested explicitly.
	_, err = s.State.AddMachine()
	c.Assert(err, IsNil)
}

func (s *AssignSuite) TestUnassignUnitFromMachineWithoutBeingAssigned(c *C) {
	// When unassigning a machine from a unit, it is possible that
	// the machine has not been previously assigned, or that it
	// was assigned but the state changed beneath us.  In either
	// case, the end state is the intended state, so we simply
	// move forward without any errors here, to avoid having to
	// handle the extra complexity of dealing with the concurrency
	// problems.
	err := s.unit.UnassignFromMachine()
	c.Assert(err, IsNil)

	// Check that the unit has no machine assigned.
	_, err = s.unit.AssignedMachineId()
	c.Assert(err, ErrorMatches, `cannot get machine id of unit "wordpress/0": unit not assigned to machine`)
}

func (s *AssignSuite) TestAssignUnitToMachineAgainFails(c *C) {
	// Check that assigning an already assigned unit to
	// a machine fails if it isn't precisely the same
	// machine.
	machineOne, err := s.State.AddMachine()
	c.Assert(err, IsNil)
	machineTwo, err := s.State.AddMachine()
	c.Assert(err, IsNil)

	err = s.unit.AssignToMachine(machineOne)
	c.Assert(err, IsNil)

	// Assigning the unit to the same machine should return no error.
	err = s.unit.AssignToMachine(machineOne)
	c.Assert(err, IsNil)

	// Assigning the unit to a different machine should fail.
	// BUG(aram): use error strings from state.
	err = s.unit.AssignToMachine(machineTwo)
	c.Assert(err, ErrorMatches, `cannot assign unit "wordpress/0" to machine 2: .*`)

	machineId, err := s.unit.AssignedMachineId()
	c.Assert(err, IsNil)
	c.Assert(machineId, Equals, 1)
}

func (s *AssignSuite) TestAssignedMachineIdWhenNotAlive(c *C) {
	machine, err := s.State.AddMachine()
	c.Assert(err, IsNil)

	err = s.unit.AssignToMachine(machine)
	c.Assert(err, IsNil)

	subCharm := s.AddTestingCharm(c, "logging")
	subSvc, err := s.State.AddService("logging", subCharm)
	c.Assert(err, IsNil)

	subUnit, err := subSvc.AddUnitSubordinateTo(s.unit)
	c.Assert(err, IsNil)

	testWhenDying(c, s.unit, noErr, noErr,
		func() error {
			_, err = s.unit.AssignedMachineId()
			return err
		},
		func() error {
			_, err = subUnit.AssignedMachineId()
			return err
		})
}

func (s *AssignSuite) TestUnassignUnitFromMachineWithChangingState(c *C) {
	// Check that unassigning while the state changes fails nicely.
	// Remove the unit for the tests.
	err := s.unit.Die()
	c.Assert(err, IsNil)
	err = s.service.RemoveUnit(s.unit)
	c.Assert(err, IsNil)

	err = s.unit.UnassignFromMachine()
	c.Assert(err, ErrorMatches, `cannot unassign unit "wordpress/0" from machine: .*`)
	_, err = s.unit.AssignedMachineId()
	c.Assert(err, ErrorMatches, `cannot get machine id of unit "wordpress/0": unit not assigned to machine`)

	err = s.service.Die()
	c.Assert(err, IsNil)
	err = s.State.RemoveService(s.service)
	c.Assert(err, IsNil)

	err = s.unit.UnassignFromMachine()
	c.Assert(err, ErrorMatches, `cannot unassign unit "wordpress/0" from machine: .*`)
	_, err = s.unit.AssignedMachineId()
	c.Assert(err, ErrorMatches, `cannot get machine id of unit "wordpress/0": unit not assigned to machine`)
}

func (s *AssignSuite) TestAssignSubordinatesToMachine(c *C) {
	// Check that assigning a principal unit assigns its subordinates too.
	subCharm := s.AddTestingCharm(c, "logging")
	logService1, err := s.State.AddService("logging1", subCharm)
	c.Assert(err, IsNil)
	logService2, err := s.State.AddService("logging2", subCharm)
	c.Assert(err, IsNil)
	log1Unit, err := logService1.AddUnitSubordinateTo(s.unit)
	c.Assert(err, IsNil)
	log2Unit, err := logService2.AddUnitSubordinateTo(s.unit)
	c.Assert(err, IsNil)

	machine, err := s.State.AddMachine()
	c.Assert(err, IsNil)
	err = log1Unit.AssignToMachine(machine)
	c.Assert(err, ErrorMatches, ".*: unit is a subordinate")
	err = s.unit.AssignToMachine(machine)
	c.Assert(err, IsNil)

	id, err := log1Unit.AssignedMachineId()
	c.Assert(err, IsNil)
	c.Check(id, Equals, machine.Id())
	id, err = log2Unit.AssignedMachineId()
	c.Check(id, Equals, machine.Id())

	// Check that unassigning the principal unassigns the
	// subordinates too.
	err = s.unit.UnassignFromMachine()
	c.Assert(err, IsNil)
	_, err = log1Unit.AssignedMachineId()
	c.Assert(err, ErrorMatches, `cannot get machine id of unit "logging1/0": unit not assigned to machine`)
	_, err = log2Unit.AssignedMachineId()
	c.Assert(err, ErrorMatches, `cannot get machine id of unit "logging2/0": unit not assigned to machine`)
}

func (s *AssignSuite) TestAssignMachineWhenDying(c *C) {
	machine, err := s.State.AddMachine()
	c.Assert(err, IsNil)

	const unitDeadErr = ".*: unit is dead"
	unit, err := s.service.AddUnit()
	c.Assert(err, IsNil)

	assignTest := func() error {
		err := unit.AssignToMachine(machine)
		err1 := unit.UnassignFromMachine()
		c.Assert(err1, IsNil)
		return err
	}
	testWhenDying(c, unit, unitDeadErr, unitDeadErr, assignTest)

	const machineDeadErr = ".*: machine is dead"
	unit, err = s.service.AddUnit()
	c.Assert(err, IsNil)
	testWhenDying(c, machine, machineDeadErr, machineDeadErr, assignTest)

	// Check that UnassignFromMachine works when the unit is dead.
	machine, err = s.State.AddMachine()
	c.Assert(err, IsNil)
	unit, err = s.service.AddUnit()
	c.Assert(err, IsNil)
	err = unit.AssignToMachine(machine)
	c.Assert(err, IsNil)
	err = unit.Die()
	c.Assert(err, IsNil)
	err = unit.UnassignFromMachine()
	c.Assert(err, IsNil)

	// Check that UnassignFromMachine works when the machine is
	// dead.
	machine, err = s.State.AddMachine()
	c.Assert(err, IsNil)
	unit, err = s.service.AddUnit()
	c.Assert(err, IsNil)
	err = unit.AssignToMachine(machine)
	c.Assert(err, IsNil)
	err = machine.Die()
	c.Assert(err, IsNil)
	err = unit.UnassignFromMachine()
	c.Assert(err, IsNil)
}

func (s *AssignSuite) TestAssignMachinePrincipalsChange(c *C) {
	machine, err := s.State.AddMachine()
	c.Assert(err, IsNil)
	unit, err := s.service.AddUnit()
	c.Assert(err, IsNil)
	err = unit.AssignToMachine(machine)
	c.Assert(err, IsNil)
	unit, err = s.service.AddUnit()
	c.Assert(err, IsNil)
	err = unit.AssignToMachine(machine)
	c.Assert(err, IsNil)
	subCharm := s.AddTestingCharm(c, "logging")
	logService, err := s.State.AddService("logging", subCharm)
	c.Assert(err, IsNil)
	_, err = logService.AddUnitSubordinateTo(unit)
	c.Assert(err, IsNil)

	doc := make(map[string][]string)
	s.ConnSuite.machines.FindId(machine.Id()).One(&doc)
	principals, ok := doc["principals"]
	if !ok {
		c.Errorf(`machine document does not have a "principals" field`)
	}
	c.Assert(principals, DeepEquals, []string{"wordpress/1", "wordpress/2"})

	err = unit.Die()
	c.Assert(err, IsNil)
	err = s.service.RemoveUnit(unit)
	c.Assert(err, IsNil)
	doc = make(map[string][]string)
	s.ConnSuite.machines.FindId(machine.Id()).One(&doc)
	principals, ok = doc["principals"]
	if !ok {
		c.Errorf(`machine document does not have a "principals" field`)
	}
	c.Assert(principals, DeepEquals, []string{"wordpress/1"})
}

func (s *AssignSuite) TestAssignUnitToUnusedMachine(c *C) {
	// Add some units to another service and allocate them to machines
	service1, err := s.State.AddService("wordpress1", s.charm)
	c.Assert(err, IsNil)
	units := make([]*state.Unit, 3)
	for i := range units {
		u, err := service1.AddUnit()
		c.Assert(err, IsNil)
		m, err := s.State.AddMachine()
		c.Assert(err, IsNil)
		err = u.AssignToMachine(m)
		c.Assert(err, IsNil)
		units[i] = u
	}
	c.Logf("added initial units")
	printMachines(c, s.State)

	// Assign the suite's unit to a machine, then remove the service
	// so the machine becomes available again.
	origMachine, err := s.State.AddMachine()
	c.Assert(err, IsNil)
	err = s.unit.AssignToMachine(origMachine)
	c.Assert(err, IsNil)
	err = s.service.Die()
	c.Assert(err, IsNil)
	// TODO s.unit.Die
	err = s.State.RemoveService(s.service)
	c.Assert(err, IsNil)
	c.Logf("removed service")
	printMachines(c, s.State)

	// Check that AssignToUnusedMachine finds the old (now unused) machine.
	newService, err := s.State.AddService("wordpress2", s.charm)
	c.Assert(err, IsNil)
	newUnit, err := newService.AddUnit()
	c.Assert(err, IsNil)
	reusedMachine, err := newUnit.AssignToUnusedMachine()
	c.Assert(err, IsNil)
	c.Assert(reusedMachine.Id(), Equals, origMachine.Id())

	c.Logf("assigned to old machine")
	printMachines(c, s.State)

	// Check that it fails when called again, even when there's an available machine
	m, err := s.State.AddMachine()
	c.Assert(err, IsNil)
	_, err = newUnit.AssignToUnusedMachine()
	c.Assert(err, ErrorMatches, `cannot assign unit "wordpress2/0" to unused machine: unit is already assigned to a machine`)
	err = m.Die()
	c.Assert(err, IsNil)
	err = s.State.RemoveMachine(m.Id())
	c.Assert(err, IsNil)

	c.Logf("failed to assign again")
	printMachines(c, s.State)

	// Try to assign another unit to an unused machine
	// and check that we can't
	newUnit, err = newService.AddUnit()
	c.Assert(err, IsNil)

	c.Logf("added unit")
	printMachines(c, s.State)
	m, err = newUnit.AssignToUnusedMachine()
	c.Assert(m, IsNil)
	c.Assert(err, ErrorMatches, `all machines in use`)

	// Add a dying machine and check that it is not chosen.
	m, err = s.State.AddMachine()
	c.Assert(err, IsNil)
	err = m.Kill()
	c.Assert(err, IsNil)
	m, err = newUnit.AssignToUnusedMachine()
	c.Assert(m, IsNil)
	c.Assert(err, ErrorMatches, `all machines in use`)
}

func printMachines(c *C, st *state.State) {
	ms, err := st.AllMachines()
	c.Assert(err, IsNil)
	for _, m := range ms {
		units, err := m.Units()
		c.Assert(err, IsNil)
		c.Logf("machine %v: %v", m, units)
	}
}

func (s *AssignSuite) TestAssignUnitToUnusedMachineWithChangingService(c *C) {
	// Check for a 'state changed' error if a service is manipulated
	// during reuse.
	err := s.service.Die()
	c.Assert(err, IsNil)
	err = s.State.RemoveService(s.service)
	c.Assert(err, IsNil)
	_, err = s.State.AddMachine()
	c.Assert(err, IsNil)

	_, err = s.unit.AssignToUnusedMachine()
	c.Assert(err, ErrorMatches, `cannot assign unit "wordpress/0" to unused machine.*: cannot get unit "wordpress/0": not found`)
}

func (s *AssignSuite) TestAssignUnitToUnusedMachineWithChangingUnit(c *C) {
	// Check for a 'state changed' error if a unit is manipulated
	// during reuse.
	err := s.unit.Die()
	c.Assert(err, IsNil)
	err = s.service.RemoveUnit(s.unit)
	c.Assert(err, IsNil)
	_, err = s.State.AddMachine()
	c.Assert(err, IsNil)

	_, err = s.unit.AssignToUnusedMachine()
	c.Assert(err, ErrorMatches, `cannot assign unit "wordpress/0" to unused machine.*: cannot get unit "wordpress/0": not found`)
}

func (s *AssignSuite) TestAssignUnitToUnusedMachineOnlyZero(c *C) {
	// Check that the unit can't be assigned to machine zero.
	_, err := s.unit.AssignToUnusedMachine()
	c.Assert(err, ErrorMatches, `all machines in use`)
}

func (s *AssignSuite) TestAssignUnitToUnusedMachineNoneAvailable(c *C) {
	// Check that assigning without unused machine fails.
	m1, err := s.State.AddMachine()
	c.Assert(err, IsNil)
	err = s.unit.AssignToMachine(m1)
	c.Assert(err, IsNil)

	newUnit, err := s.service.AddUnit()
	c.Assert(err, IsNil)

	_, err = newUnit.AssignToUnusedMachine()
	c.Assert(err, ErrorMatches, `all machines in use`)
}
