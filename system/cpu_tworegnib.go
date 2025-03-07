package system

import (
	"fmt"

	"github.com/ShyProton/chip8/system/ops"
)

// DRW is the only TwoRegNib instruction.
func (sys *System) tryRunIfTwoRegNib(inst ops.Instruction) (bool, error) {
	twoRegNibInst := inst.ApplyOpcodeMask(ops.TwoRegNib)
	x, y, n := inst.GetTwoRegNib()

	if twoRegNibInst != ops.DRW {
		return false, nil
	}

	drawX, drawY := int(sys.registers.V[x]), int(sys.registers.V[y])
	erasure := false

	for i := range int(n) {
		sprRow, err := sys.memory.ByteAt(int(sys.registers.I) + i)
		if err != nil {
			return true, fmt.Errorf("could not fully access sprite for DRW operation:\n%w", err)
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
