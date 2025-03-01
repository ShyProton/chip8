package system

import (
	"fmt"

	"github.com/ShyProton/chip8/system/memory"
)

type System struct {
	memory    memory.Memory
	registers Registers
	stack     CallStack
	io        IO
}

type instExecutionError struct {
	inst Instruction
	err  error
}

func (err instExecutionError) Error() string {
	return fmt.Sprintf("error encountered trying to execute instruction '%04X':\n%v", err.inst, err.err)
}

func NewSystem(romPath string) (*System, error) {
	system := new(System)

	system.memory.LoadFont()

	if err := system.memory.LoadRom(romPath); err != nil {
		return nil, fmt.Errorf("error while loading rom into memory: %w", err)
	}

	if err := system.io.Init(romPath); err != nil {
		return nil, fmt.Errorf("error while initializing the system graphics: %w", err)
	}

	return system, nil
}

func (sys *System) Run() error {
	for ; ; sys.memory.IncPC() {
		even, odd := sys.memory.GetInstBytes()
		inst := InstFromBytes(even, odd)

		if err := sys.Execute(inst); err != nil {
			return instExecutionError{inst, err}
		}
	}
}

// Stuff that the system needs to wrap up before the program exits.
func (sys *System) Exit() {
	sys.io.graphics.Finish()
}
