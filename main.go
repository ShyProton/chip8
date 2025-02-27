package main

import (
	"fmt"
	"os"
)

type System struct {
	memory    Memory
	registers Registers
	stack     CallStack
	display   Display
	// TODO: Display - for handing the visual output.
	// TODO: CPU - for handling the decoding and execution of instructions.
}

type InstExecutionError struct {
	inst Instruction
	err  error
}

func (err InstExecutionError) Error() string {
	return fmt.Sprintf("error encountered trying to execute instruction '%04X':\n%v", err.inst, err.err)
}

func (sys *System) Run() error {
	err := sys.display.Init()
	if err != nil {
		return err
	}
	defer sys.display.Finish()

	sys.registers.PC = RomStart // Starts at address 512.

	for ; ; sys.registers.IncProgramCounter() {
		even, odd := sys.memory.GetInstBytes(&sys.registers)
		inst := InstFromBytes(even, odd)

		err = sys.Execute(inst)
		if err != nil {
			return InstExecutionError{inst, err}
		}

		fmt.Printf("Dual-Byte: %X\n", inst)
	}
}

func main() {
	var system System

	err := system.memory.LoadRom("rom.ch8")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)

		return
	}

	err = system.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)

		return
	}
}
