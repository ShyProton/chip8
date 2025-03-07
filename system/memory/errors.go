package memory

import "fmt"

type romTooLargeError struct {
	RomSize int
}

func (err romTooLargeError) Error() string {
	return fmt.Sprintf("rom file size %d is too large (max %d)",
		err.RomSize, romCapacity)
}

type outOfBoundsError struct {
	Address int
}

func (err outOfBoundsError) Error() string {
	return fmt.Sprintf("could not access memory contents at address %04d (max %d)",
		err.Address, memoryCapacity-1)
}

type invalidFontAccessError struct {
	digit byte
}

func (err invalidFontAccessError) Error() string {
	return fmt.Sprintf("a font character for value '%X' does not exist, "+
		"as it is larger than a single digit hexadecimal value (%X)",
		err.digit, fontChars-1)
}

type invalidPCAssignmentError struct {
	PC uint16
}

func (err invalidPCAssignmentError) Error() string {
	return fmt.Sprintf("could not set program counter to '%d', "+
		"as it must be less than %d and must be even",
		err.PC, memoryCapacity-1)
}
