package main

import (
	"fmt"
	"os"
)

type System struct {
	memory    Memory
	registers Registers
	stack     CallStack
	io        IO
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

func main() {
	var system System

	system.memory.LoadFont()

	romName := "2-ibm-logo.ch8"

	err := system.memory.LoadRom(romName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)

		return
	}

	err = system.io.Init(romName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)

		return
	}
	defer system.io.graphics.Finish()

	err = system.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)

		return
	}
}
