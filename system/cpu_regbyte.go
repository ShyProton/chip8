package system

import (
	"math/rand/v2"

	"github.com/ShyProton/chip8/system/ops"
)

// RegByte instructions include:
// SE, SNE, LD, ADD, RND.
func (sys *System) tryRunIfRegByte(inst ops.Instruction) (bool, error) {
	regByteInst := inst.ApplyOpcodeMask(ops.RegByte)
	x, b := inst.GetRegByte()

	var err error

	switch regByteInst {
	case ops.SE: // Skip next instruction if Vx == byte.
		if sys.registers.V[x] == b {
			sys.memory.IncPC()
		}
	case ops.SNE: // Skip next instruction if Vx != byte.
		if sys.registers.V[x] != b {
			sys.memory.IncPC()
		}
	case ops.LD: // Set Vx = byte.
		sys.registers.V[x] = b
	case ops.ADD: // Set Vx = Vx + byte.
		sys.registers.V[x] += b
	case ops.RND: // Set Vx = random byte AND passed byte.
		sys.registers.V[x] = byte(rand.IntN(256)) & b
	default:
		return false, nil
	}

	return true, err
}
