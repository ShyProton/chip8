package main

import "fmt"

type InstTypeRunner = func(Instruction) (bool, error)

func (sys *System) Execute(inst Instruction) error {
	instructionTypeRunners := [...]InstTypeRunner{sys.tryRunIfExact, sys.tryRunIfAddr}

	for _, tryRunInstruction := range instructionTypeRunners {
		found, err := tryRunInstruction(inst)
		if err != nil {
			return err
		}

		if found {
			return nil
		}
	}

	fmt.Printf("Instruction %X not a valid command, ignoring...\n", inst)
	return nil
}

// Exact instructions include CLS, RET
func (sys *System) tryRunIfExact(inst Instruction) (bool, error) {
	exactInst := inst.ApplyOpcodeMask(Exact)

	found := false
	var err error

	switch exactInst {
	case CLS:
		sys.CLS()
		found = true
	case RET:
		err = sys.RET()
		found = true
	}

	return found, err
}

func (sys *System) tryRunIfAddr(inst Instruction) (bool, error) {
	addrInst := inst.ApplyOpcodeMask(Addr)
	address := inst.GetAddr()

	found := false
	var err error

	switch addrInst {
	case JP:
		sys.JP(address)
		found = true
	case CALL:
		err = sys.CALL(address)
		found = true
	}

	return found, err
}

// Clear the display.
func (sys *System) CLS() {
	// TODO: Must clear the system's display.
}

// The interpreter sets the program counter to the address at the top of the
// stack, then subtracts 1 from the stack pointer.
func (sys *System) RET() error {
	err := sys.stack.Pop(&sys.registers)
	return err
}

// Jump to location nnn.
func (sys *System) JP(nnn uint16) {
	sys.registers.PC = nnn
}

// Call subroutine at nnn.
func (sys *System) CALL(nnn uint16) error {
	err := sys.stack.Push(&sys.registers)
	sys.registers.PC = nnn
	return err
}
