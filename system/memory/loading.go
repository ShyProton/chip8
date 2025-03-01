package memory

import (
	"fmt"
	"os"
)

const romCapacity = memoryCapacity - romStart // How large a program can be.

const (
	fontChars    = 16 // Number of font characters within the interpreter.
	fontCharRows = 5  // How many bytes the characters are composed of.
)

func (mem *Memory) LoadRom(romPath string) error {
	rom, err := os.ReadFile(romPath)
	if err != nil {
		return fmt.Errorf("could not read rom file: %w", err)
	}

	if len(rom) > romCapacity {
		return romTooLargeError{len(rom)}
	}

	for i, romByte := range rom {
		mem.ram[i+romStart] = romByte
	}

	mem.pc = romStart
	mem.nextpc = mem.pc + 2

	return nil
}

// See http://devernay.free.fr/hacks/chip8/C8TECH10.HTM#font,
// which outlines the internal font definitions.
func (mem *Memory) LoadFont() {
	digits := [fontChars][fontCharRows]byte{
		{0xF0, 0x90, 0x90, 0x90, 0xF0}, // Zero.     (0)
		{0x20, 0x60, 0x20, 0x20, 0x70}, // One.      (1)
		{0xF0, 0x10, 0xF0, 0x80, 0xF0}, // Two.      (2)
		{0xF0, 0x10, 0xF0, 0x10, 0xF0}, // Three.    (3)
		{0x90, 0x90, 0xF0, 0x10, 0x10}, // Four.     (4)
		{0xF0, 0x80, 0xF0, 0x10, 0xF0}, // Five.     (5)
		{0xF0, 0x80, 0xF0, 0x90, 0xF0}, // Six.      (6)
		{0xF0, 0x10, 0x20, 0x40, 0x40}, // Seven.    (7)
		{0xF0, 0x90, 0xF0, 0x90, 0xF0}, // Eight.    (8)
		{0xF0, 0x90, 0xF0, 0x10, 0xF0}, // Nine.     (9)
		{0xF0, 0x90, 0xF0, 0x90, 0x90}, // Ten.      (A)
		{0xE0, 0x90, 0xE0, 0x90, 0xE0}, // Eleven.   (B)
		{0xF0, 0x80, 0x80, 0x80, 0xF0}, // Twelve.   (C)
		{0xE0, 0x90, 0x90, 0x90, 0xE0}, // Thirteen. (D)
		{0xF0, 0x80, 0xF0, 0x80, 0xF0}, // Fourteen. (E)
		{0xF0, 0x80, 0xF0, 0x80, 0x80}, // Fifteen.  (F)
	}

	for i, digit := range digits {
		for j, row := range digit {
			mem.ram[i*len(digit)+j] = row
		}
	}
}

func (mem *Memory) LoadFromBytes(addr int, bytes []byte) error {
	// Cannot load registers into memory if we exceed the memory capacity while doing so.
	largestAddr := addr + len(bytes) - 1
	if largestAddr >= memoryCapacity {
		return outOfBoundsError{memoryCapacity}
	}

	for i, b := range bytes {
		mem.ram[addr+i] = b
	}

	return nil
}
