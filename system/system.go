package system

import "fmt"

type System struct {
	memory    Memory
	registers Registers
	stack     CallStack
	io        IO
}

func NewSystem(romPath string) (*System, error) {
	system := new(System)

	system.memory.LoadFont()

	err := system.memory.LoadRom(romPath)
	if err != nil {
		return nil, fmt.Errorf("error while loading rom into memory: %w", err)
	}

	err = system.io.Init(romPath)
	if err != nil {
		return nil, fmt.Errorf("error while initializing the system graphics: %w", err)
	}

	return system, nil
}

type InstExecutionError struct {
	inst Instruction
	err  error
}

func (err InstExecutionError) Error() string {
	return fmt.Sprintf("error encountered trying to execute instruction '%04X':\n%v", err.inst, err.err)
}

func (sys *System) Run() error {
	sys.registers.PC = RomStart // Starts at address 512.

	for ; ; sys.registers.IncProgramCounter() {
		even, odd := sys.memory.GetInstBytes(&sys.registers)
		inst := InstFromBytes(even, odd)

		if err := sys.Execute(inst); err != nil {
			return InstExecutionError{inst, err}
		}
	}
}

// Stuff that the system needs to wrap up before the program exits.
func (sys *System) Exit() {
	sys.io.graphics.Finish()
}
