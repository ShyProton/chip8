package main

import (
	"fmt"
	"os"
)

const (
	MemoryCapacity = 4_096                     // 4KB of memory.
	RomStart       = 512                       // Starting address of Chip-8 programs.
	RomCapacity    = MemoryCapacity - RomStart // How large a program can be.
)

const (
	FontChars    = 16
	FontCharRows = 5
)

type Memory [MemoryCapacity]byte

type RomTooLargeError struct {
	RomSize int
}

func (err RomTooLargeError) Error() string {
	return fmt.Sprintf("rom file size %d is too large (max %d)", err.RomSize, RomCapacity)
}

func (mem *Memory) LoadRom(filePath string) error {
	rom, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("could not read rom file: %w", err)
	}

	if len(rom) > RomCapacity {
		return RomTooLargeError{len(rom)}
	}

	for i, romByte := range rom {
		mem[i+RomStart] = romByte
	}

	return nil
}

// See http://devernay.free.fr/hacks/chip8/C8TECH10.HTM#font,
// which outlines the internal font definitions.
func (mem *Memory) LoadFont() {
	digits := [FontChars][FontCharRows]byte{
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
			mem[i*len(digit)+j] = row
		}
	}
}

func (mem *Memory) GetInstBytes(reg *Registers) (byte, byte) {
	return mem[reg.PC], mem[reg.PC+1]
}

type OutOfBoundsError struct {
	Address int
}

func (err OutOfBoundsError) Error() string {
	return fmt.Sprintf("could not access memory contents at address %d (max %d)", err.Address, MemoryCapacity-1)
}

func (mem *Memory) ByteAt(i int) (*byte, error) {
	if i >= MemoryCapacity {
		return nil, OutOfBoundsError{i}
	}

	return &mem[i], nil
}
