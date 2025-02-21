package main

// func (sys *System) Execute(inst Instruction) {
// 	switch {
// 	case inst.MatchesOpcode(sys.opcodes[CLS]):
// 		// TODO: CLS
// 	case inst.MatchesOpcode(sys.opcodes[RET]):
// 		// TODO: ADD
// 	case inst.MatchesOpcode(sys.opcodes[JP]):
// 		// TODO: JP
// 	case inst.MatchesOpcode(sys.opcodes[CALL]):
// 		// TODO: CALL
// 	case inst.MatchesOpcode(sys.opcodes[SE]):
// 		// TODO: SE
// 	}
// 	// ...
// }

// TODO: Possible different interfaces for the different types of masks.
type ExactInst interface {
	Run() error
}

type AddrInst interface {
	Run(addr uint16) error
}

type RegInst interface {
	Run(reg uint) error
}

func (sys *System) Execute(inst Instruction) {
	exactInst := inst.ApplyOpcodeMask(Exact)
	switch exactInst {
	case sys.opcodes[CLS].Code:
		// TODO: CLS
	case sys.opcodes[RET].Code:
		// TODO: RET
	}
}
