package system

import (
	"math"
	"math/rand/v2"
)

type InstTypeRunner = func(Instruction) (bool, error)

func (sys *System) Execute(inst Instruction) error {
	instructionTypeRunners := [...]InstTypeRunner{
		sys.tryRunIfExact,
		sys.tryRunIfAddr,
		sys.tryRunIfRegByte,
		sys.tryRunIfTwoReg,
		sys.tryRunIfReg,
		sys.tryRunIfTwoRegNib,
	}

	for _, tryRunInstruction := range instructionTypeRunners {
		found, err := tryRunInstruction(inst)
		if err != nil {
			return err
		}

		if found {
			return nil
		}
	}

	return nil
}

// Exact instructions include CLS and RET.
func (sys *System) tryRunIfExact(inst Instruction) (bool, error) {
	exactInst := inst.ApplyOpcodeMask(Exact)

	var err error

	switch exactInst {
	case CLS: // Clear the display.
		sys.io.graphics.Buf.Clear()
	case RET: // Return from a subroutine.
		err = sys.memory.PopCallStack()
	default:
		return false, nil
	}

	return true, err
}

// Address instructions include JP and CALL.
func (sys *System) tryRunIfAddr(inst Instruction) (bool, error) {
	addrInst := inst.ApplyOpcodeMask(Addr)
	address := inst.GetAddr()

	var err error

	switch addrInst {
	case JP: // Jump to location at address.
		err = sys.memory.QueueNextPC(address)
	case CALL: // Call subroutine at address.
		err = sys.memory.PushCallStack()
		if err != nil {
			break
		}

		err = sys.memory.QueueNextPC(address)
	case LDI: // Set I = nnn.
		sys.registers.I = address
	case JPV: // Jump to location nnn + V0.
		err = sys.memory.QueueNextPC(address + uint16(sys.registers.V[0]))
	default:
		return false, nil
	}

	return true, err
}

// RegByte instructions include:
// SE, SNE, LD, ADD, RND.
func (sys *System) tryRunIfRegByte(inst Instruction) (bool, error) {
	regByteInst := inst.ApplyOpcodeMask(RegByte)
	x, b := inst.GetRegByte()

	var err error

	switch regByteInst {
	case SE: // Skip next instruction if Vx == byte.
		if sys.registers.V[x] == b {
			sys.memory.IncPC()
		}
	case SNE: // Skip next instruction if Vx != byte.
		if sys.registers.V[x] != b {
			sys.memory.IncPC()
		}
	case LD: // Set Vx = byte.
		sys.registers.V[x] = b
	case ADD: // Set Vx = Vx + byte.
		sys.registers.V[x] += b
	case RND: // Set Vx = random byte AND passed byte.
		sys.registers.V[x] = byte(rand.IntN(256)) & b
	default:
		return false, nil
	}

	return true, err
}

// TwoReg instructions include:
// RegLD, RegOR, RegAND, RegXOR, RegADD, RegSUB, RegSHR, RegSUBN, RegSHL, RegSNE.
func (sys *System) tryRunIfTwoReg(inst Instruction) (bool, error) {
	twoRegInst := inst.ApplyOpcodeMask(TwoReg)
	x, y, _ := inst.GetTwoRegNib()

	switch twoRegInst {
	case RegLD: // Set Vx = Vy.
		sys.registers.V[x] = sys.registers.V[y]
	case RegOR: // Set Vx = Vx OR Vy.
		sys.registers.V[x] |= sys.registers.V[y]
	case RegAND: // Set Vx = Vx AND Vy.
		sys.registers.V[x] &= sys.registers.V[y]
	case RegXOR: // Set Vx = Vx XOR Vy.
		sys.registers.V[x] ^= sys.registers.V[y]
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
	case RegSE:
		if sys.registers.V[x] == sys.registers.V[y] {
			sys.memory.IncPC()
		}
	case RegSNE: // Skip next instruction if Vx != Vy.
		if sys.registers.V[x] != sys.registers.V[y] {
			sys.memory.IncPC()
		}
	default:
		return false, nil
	}

	return true, nil
}

// Reg instructions include:
// SKP, SKPNP, LDDT, LDK, DTLD, STLD, ADDI, LDF, LDB, LDV, VLD.
func (sys *System) tryRunIfReg(inst Instruction) (bool, error) {
	regInst := inst.ApplyOpcodeMask(Reg)
	x, _ := inst.GetRegByte()

	var err error

	switch regInst {
	case SKP: // TODO: Skip next instruction if the key with the value of Vx is pressed.
	case SKPNP: // TODO: Skip next instruction if the key with the value of Vx is not pressed.
	case LDDT: // Set Vx = delay timer value.
		sys.registers.V[x] = sys.registers.DT
	case LDK: // Wait for a key press, store the value of the key in Vx.
	case DTLD: // Set delay timer = Vx.
		sys.registers.DT = sys.registers.V[x]
	case STLD: // Set sound timer = Vx.
		sys.registers.ST = sys.registers.V[x]
	case ADDI: // Set I = I + Vx.
		sys.registers.I += uint16(sys.registers.V[x])
	case LDF: // Set I = location of sprite for digit Vx.
		sys.registers.I, err = sys.memory.FontAddr(sys.registers.V[x])
	case LDB: // Store BCD representation of Vx in memory locations I, I+1, and I+2.
		hundreds := sys.registers.V[x] / 100
		tens := (sys.registers.V[x] / 10) % 10
		ones := sys.registers.V[x] % 10

		err = sys.memory.LoadFromBytes(int(sys.registers.I), []byte{hundreds, tens, ones})
	case LDV: // Store registers V0 through Vx in memory starting at location I.
		err = sys.memory.LoadFromBytes(int(sys.registers.I), sys.registers.V[:x+1])
	case VLD: // Read registers V0 through Vx from memory starting at location I.
		err = sys.memory.ReadToBytes(int(sys.registers.I), sys.registers.V[:x+1])
	default:
		return false, nil
	}

	return true, err
}

// DRW is the only TwoRegNib instruction.
func (sys *System) tryRunIfTwoRegNib(inst Instruction) (bool, error) {
	twoRegNibInst := inst.ApplyOpcodeMask(TwoRegNib)
	x, y, n := inst.GetTwoRegNib()

	if twoRegNibInst != DRW {
		return false, nil
	}

	drawX, drawY := int(sys.registers.V[x]), int(sys.registers.V[y])
	erasure := false

	for i := range int(n) {
		sprRow, err := sys.memory.ByteAt(int(sys.registers.I) + i)
		if err != nil {
			return true, err
		}

		erasure = erasure || sys.io.DrawRow(drawX, drawY+i, *sprRow)
	}

	if erasure {
		sys.registers.V[0xF] = 1
	} else {
		sys.registers.V[0xF] = 0
	}

	sys.io.graphics.Show()

	return true, nil
}
