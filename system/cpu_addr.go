package system

import (
	"fmt"

	"github.com/ShyProton/chip8/system/ops"
)

// Address instructions include JP and CALL.
func (sys *System) tryRunIfAddr(inst ops.Instruction) (bool, error) {
	addrInst := inst.ApplyOpcodeMask(ops.Addr)
	addr := inst.GetAddr()

	var err error

	switch addrInst {
	case ops.JP: // Jump to location at address.
		err = sys.memory.QueueNextPC(addr)
	case ops.CALL: // Call subroutine at address.
		if err := sys.memory.PushCallStack(); err != nil {
			break
		}

		err = sys.memory.QueueNextPC(addr)
	case ops.LDI: // Set I = nnn.
		sys.registers.I = addr
	case ops.JPV: // Jump to location nnn + V0.
		err = sys.memory.QueueNextPC(addr + uint16(sys.registers.V[0]))
	default:
		return false, nil
	}

	if err != nil {
		err = fmt.Errorf("addr instruction error:\n%w", err)
	}

	return true, err
}
