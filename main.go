package main

import (
	"fmt"
	"os"
)

type System struct {
	memory    Memory
	registers Registers
	stack     CallStack
	// TODO: Display - for handing the visual output.
	// TODO: CPU - for handling the decoding and execution of instructions.
}

func (sys *System) Run() error {
	sys.registers.PC = RomStart // Starts at address 512.

	for ; ; sys.registers.PC += 2 {
		even, odd, err := sys.memory.GetInstBytes(&sys.registers)
		if err != nil {
			return err
		}

		inst := InstFromBytes(even, odd)

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
