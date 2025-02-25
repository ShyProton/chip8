package main

import (
	"fmt"
	"math/rand"
)

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

// Exact instructions include CLS and RET
func (sys *System) tryRunIfExact(inst Instruction) (bool, error) {
	exactInst := inst.ApplyOpcodeMask(Exact)

	var err error

	switch exactInst {
	case CLS: // Clear the display.
		// TODO: When display is implemented.
	case RET: // Return from a subroutine.
		err = sys.stack.Pop(&sys.registers)
	default:
		return false, err
	}

	return true, err
}

// Address instructions include JP and CALL
func (sys *System) tryRunIfAddr(inst Instruction) (bool, error) {
	addrInst := inst.ApplyOpcodeMask(Addr)
	address := inst.GetAddr()

	var err error

	switch addrInst {
	case JP: // Jump to location at address.
		sys.registers.PC = address
	case CALL: // Call subroutine at address.
		err = sys.stack.Push(&sys.registers)
		sys.registers.PC = address
	default:
		return false, err
	}

	return true, err
}

// RegByte instructions include:
// SE, SNE, LD, ADD, RND
func (sys *System) tryRunIfRegByte(inst Instruction) (bool, error) {
	regByteInst := inst.ApplyOpcodeMask(RegByte)
	reg, b := inst.GetRegByte()

	var err error

	switch regByteInst {
	case SE: // Skip next instruction if Vx == byte.
		if sys.registers.V[reg] == b {
			sys.registers.IncProgramCounter()
		}
	case SNE: // Skip next instruction if Vx != byte.
		if sys.registers.V[reg] != b {
			sys.registers.IncProgramCounter()
		}
	case LD: // Set Vx = byte.
		sys.registers.V[reg] = b
	case ADD: // Set Vx = Vx + byte.
		sys.registers.V[reg] += b
	case RND: // Set Vx = random byte AND passed byte.
		sys.registers.V[reg] = byte(rand.Intn(256)) & b
	default:
		return false, err
	}

	return true, err
}
