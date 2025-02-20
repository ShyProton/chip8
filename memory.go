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

type Memory [MemoryCapacity]byte

type RomTooLargeError struct {
	RomSize int
}

func (err RomTooLargeError) Error() string {
	return fmt.Sprintf("rom file size %d is too large (max %d)", err.RomSize, RomCapacity)
}

type OutOfBoundsError struct {
	Address uint16
}

func (err OutOfBoundsError) Error() string {
	return fmt.Sprintf("could not access memory contents at address %d (max %d)", err.Address, MemoryCapacity-InstBytes)
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

func (mem *Memory) GetInstBytes(reg *Registers) (byte, byte, error) {
	if reg.PC > MemoryCapacity-InstBytes {
		return 0, 0, OutOfBoundsError{reg.PC}
	}

	return mem[reg.PC], mem[reg.PC+1], nil
}
