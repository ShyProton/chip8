package system

import "github.com/ShyProton/chip8/system/ops"

type InstTypeRunner = func(ops.Instruction) (bool, error)

func (sys *System) Execute(inst ops.Instruction) error {
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
