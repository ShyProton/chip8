package system

import "github.com/ShyProton/chip8/system/ops"

// Exact commands include CLS and RET.
func (sys *System) tryRunIfExact(inst ops.Instruction) (bool, error) {
	exactInst := inst.ApplyOpcodeMask(ops.Exact)

	var err error

	switch exactInst {
	case ops.CLS: // Clear the display.
		sys.io.graphics.Buf.Clear()
	case ops.RET: // Return from a subroutine.
		err = sys.memory.PopCallStack()
	default:
		return false, nil
	}

	return true, err
}
