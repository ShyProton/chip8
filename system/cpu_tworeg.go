package system

import (
	"math"

	"github.com/ShyProton/chip8/system/ops"
)

// TwoReg instructions include:
// RegLD, RegOR, RegAND, RegXOR, RegADD, RegSUB, RegSHR, RegSUBN, RegSHL, RegSNE.
func (sys *System) tryRunIfTwoReg(inst ops.Instruction) (bool, error) {
	twoRegInst := inst.ApplyOpcodeMask(ops.TwoReg)
	x, y, _ := inst.GetTwoRegNib()

	switch twoRegInst {
	case ops.RegLD: // Set Vx = Vy.
		sys.registers.V[x] = sys.registers.V[y]
	case ops.RegOR: // Set Vx = Vx OR Vy.
		sys.registers.V[x] |= sys.registers.V[y]
	case ops.RegAND: // Set Vx = Vx AND Vy.
		sys.registers.V[x] &= sys.registers.V[y]
	case ops.RegXOR: // Set Vx = Vx XOR Vy.
		sys.registers.V[x] ^= sys.registers.V[y]
	case ops.RegADD: // Set Vx = Vx + Vy, set VF = carry.
		if uint(sys.registers.V[x])+uint(sys.registers.V[y]) > math.MaxUint8 {
			sys.registers.V[0xF] = 1
		}

		sys.registers.V[x] += sys.registers.V[y]
	case ops.RegSUB: // Set Vx = Vx - Vy, set VF = NOT borrow.
		if sys.registers.V[x] > sys.registers.V[y] {
			sys.registers.V[0xF] = 1
		}

		sys.registers.V[x] -= sys.registers.V[y]
	case ops.RegSHR: // Set Vx = Vx SHR 1.
		if sys.registers.V[x]%2 == 1 {
			sys.registers.V[0xF] = 1
		}

		sys.registers.V[x] /= 2
	case ops.RegSUBN: // Set Vx = Vy - Vx, set VF = NOT borrow.
		if sys.registers.V[y] > sys.registers.V[x] {
			sys.registers.V[0xF] = 1
		}

		sys.registers.V[x] = sys.registers.V[y] - sys.registers.V[x]
	case ops.RegSHL: // Set Vx = Vx SHL 1.
		if sys.registers.V[x]%2 == 1 {
			sys.registers.V[0xF] = 1
		}

		sys.registers.V[x] *= 2
	case ops.RegSE:
		if sys.registers.V[x] == sys.registers.V[y] {
			sys.memory.IncPC()
		}
	case ops.RegSNE: // Skip next instruction if Vx != Vy.
		if sys.registers.V[x] != sys.registers.V[y] {
			sys.memory.IncPC()
		}
	default:
		return false, nil
	}

	return true, nil
}
