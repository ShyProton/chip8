package system

import (
	"fmt"

	"github.com/ShyProton/chip8/system/memory"
	"github.com/ShyProton/chip8/system/ops"
)

const RegisterCount = 16

type System struct {
	memory    memory.Memory // System RAM.
	io        IO            // Handles display output and keyboard input.
	registers struct {
		V [RegisterCount]byte // 16 General-purpose 8-bit registers.
		I uint16              // Generally used to store memory addresses.

		// Pseudo-registers, not accessible from Chip-8 programs.
		// These should prooobably go somewhere else, belonging to each subsystem.
		SP byte // Stack Pointer.
		ST byte // Sound Timer.
		DT byte // Delay Timer.
	}
}

type instExecutionError struct {
	inst ops.Instruction
	err  error
}

func (err instExecutionError) Error() string {
	return fmt.Sprintf("error encountered trying to execute instruction '%04X':\n%v", err.inst, err.err)
}

func NewSystem(romPath string) (*System, error) {
	system := new(System)

	system.memory.LoadFont()

	if err := system.memory.LoadRom(romPath); err != nil {
		return nil, fmt.Errorf("error while loading rom into memory:\n%w", err)
	}

	if err := system.io.Init(romPath); err != nil {
		return nil, fmt.Errorf("error while initializing the system graphics:\n%w", err)
	}

	return system, nil
}

func (sys *System) Run() error {
	for ; ; sys.memory.IncPC() {
		even, odd := sys.memory.GetInstBytes()
		inst := ops.InstFromBytes(even, odd)

		if err := sys.execute(inst); err != nil {
			return instExecutionError{inst, err}
		}
	}
}

// Stuff that the system needs to wrap up before the program exits.
func (sys *System) Exit() {
	sys.io.graphics.Finish()
}
