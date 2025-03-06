package system

import "github.com/ShyProton/chip8/system/ops"

// Reg instructions include:
// SKP, SKPNP, LDDT, LDK, DTLD, STLD, ADDI, LDF, LDB, LDV, VLD.
func (sys *System) tryRunIfReg(inst ops.Instruction) (bool, error) {
	regInst := inst.ApplyOpcodeMask(ops.Reg)
	x, _ := inst.GetRegByte()

	var err error

	switch regInst {
	case ops.SKP: // TODO: Skip next instruction if the key with the value of Vx is pressed.
	case ops.SKPNP: // TODO: Skip next instruction if the key with the value of Vx is not pressed.
	case ops.LDDT: // Set Vx = delay timer value.
		sys.registers.V[x] = sys.registers.DT
	case ops.LDK: // Wait for a key press, store the value of the key in Vx.
	case ops.DTLD: // Set delay timer = Vx.
		sys.registers.DT = sys.registers.V[x]
	case ops.STLD: // Set sound timer = Vx.
		sys.registers.ST = sys.registers.V[x]
	case ops.ADDI: // Set I = I + Vx.
		sys.registers.I += uint16(sys.registers.V[x])
	case ops.LDF: // Set I = location of sprite for digit Vx.
		sys.registers.I, err = sys.memory.FontAddr(sys.registers.V[x])
	case ops.LDB: // Store BCD representation of Vx in memory locations I, I+1, and I+2.
		hundreds := sys.registers.V[x] / 100
		tens := (sys.registers.V[x] / 10) % 10
		ones := sys.registers.V[x] % 10

		err = sys.memory.LoadFromBytes(int(sys.registers.I), []byte{hundreds, tens, ones})
	case ops.LDV: // Store registers V0 through Vx in memory starting at location I.
		err = sys.memory.LoadFromBytes(int(sys.registers.I), sys.registers.V[:x+1])
	case ops.VLD: // Read registers V0 through Vx from memory starting at location I.
		err = sys.memory.ReadToBytes(int(sys.registers.I), sys.registers.V[:x+1])
	default:
		return false, nil
	}

	return true, err
}
