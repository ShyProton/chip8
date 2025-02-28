package system

const RegisterCount = 16

type Registers struct {
	V [RegisterCount]byte // 16 General-purpose 8-bit registers.
	I uint16              // Generally used to store memory addresses.

	// Pseudo-registers, not accessible from Chip-8 programs.
	// These should prooobably go somewhere else, belonging to each subsystem.
	PC uint16 // Program Counter.
	SP byte   // Stack Pointer.
	ST byte   // Sound Timer.
	DT byte   // Delay Timer.
}
