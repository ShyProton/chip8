package main

import (
	"fmt"
	"math"
	"math/rand"
)

type InstTypeRunner = func(Instruction) (bool, error)

func (sys *System) Execute(inst Instruction) error {
	instructionTypeRunners := [...]InstTypeRunner{sys.tryRunIfExact, sys.tryRunIfAddr, sys.tryRunIfRegByte, sys.tryRunIfTwoReg}

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
	x, b := inst.GetRegByte()

	var err error

	switch regByteInst {
	case SE: // Skip next instruction if Vx == byte.
		if sys.registers.V[x] == b {
			sys.registers.IncProgramCounter()
		}
	case SNE: // Skip next instruction if Vx != byte.
		if sys.registers.V[x] != b {
			sys.registers.IncProgramCounter()
		}
	case LD: // Set Vx = byte.
		sys.registers.V[x] = b
	case ADD: // Set Vx = Vx + byte.
		sys.registers.V[x] += b
	case RND: // Set Vx = random byte AND passed byte.
		sys.registers.V[x] = byte(rand.Intn(256)) & b
	default:
		return false, err
	}

	return true, err
}

// TwoReg instructions include:
// RegLD, RegOR, RegAND, RegXOR, RegADD, RegSUB, RegSHR, RegSUBN, RegSHL, RegSNE
func (sys *System) tryRunIfTwoReg(inst Instruction) (bool, error) {
	twoRegInst := inst.ApplyOpcodeMask(TwoReg)
	x, y := inst.GetTwoReg()

	switch twoRegInst {
	case RegLD: // Set Vx = Vy.
		sys.registers.V[x] = sys.registers.V[y]
	case RegOR: // Set Vx = Vx OR Vy.
		sys.registers.V[x] = sys.registers.V[x] | sys.registers.V[y]
	case RegAND: // Set Vx = Vx AND Vy.
		sys.registers.V[x] = sys.registers.V[x] & sys.registers.V[y]
	case RegXOR: // Set Vx = Vx XOR Vy.
		sys.registers.V[x] = sys.registers.V[x] ^ sys.registers.V[y]
	case RegADD: // Set Vx = Vx + Vy, set VF = carry.
		if uint(sys.registers.V[x])+uint(sys.registers.V[y]) > math.MaxUint8 {
			sys.registers.V[0xF] = 1
		}

		sys.registers.V[x] += sys.registers.V[y]
	case RegSUB: // Set Vx = Vx - Vy, set VF = NOT borrow.
		if sys.registers.V[x] > sys.registers.V[y] {
			sys.registers.V[0xF] = 1
		}

		sys.registers.V[x] -= sys.registers.V[y]
	case RegSHR: // Set Vx = Vx SHR 1.
		if sys.registers.V[x]%2 == 1 {
			sys.registers.V[0xF] = 1
		}

		sys.registers.V[x] /= 2
	case RegSUBN: // Set Vx = Vy - Vx, set VF = NOT borrow.
		if sys.registers.V[y] > sys.registers.V[x] {
			sys.registers.V[0xF] = 1
		}

		sys.registers.V[x] = sys.registers.V[y] - sys.registers.V[x]
	case RegSHL: // Set Vx = Vx SHL 1.
		if sys.registers.V[x]%2 == 1 {
			sys.registers.V[0xF] = 1
		}

		sys.registers.V[x] *= 2
	case RegSNE: // Skip next instruction if Vx != Vy.
		if sys.registers.V[x] != sys.registers.V[y] {
			sys.registers.IncProgramCounter()
		}
	default:
		return false, nil
	}

	return true, nil
}
