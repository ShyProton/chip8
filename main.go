package main

import (
	"fmt"
	"os"

	"github.com/ShyProton/chip8/system"
)

func main() {
	romPath := "roms/2-ibm-logo.ch8"

	sys, err := system.NewSystem(romPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)

		return
	}
	defer sys.Exit()

	err = sys.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)

		return
	}
}
