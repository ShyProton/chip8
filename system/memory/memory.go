package memory

import "github.com/ShyProton/chip8/system/ops"

const (
	memoryCapacity = 4_096 // 4KB of memory.
	romStart       = 512   // Starting address of Chip-8 programs.
)

type Memory struct {
	ram        [memoryCapacity]byte // System's Random Access Memory.
	cstack     callStack            // Routine call stack for storing addresses to return to.
	pc, nextpc uint16               // Program Counter indexes memory contents to get instructions.
}

func (mem *Memory) QueueNextPC(nextpc uint16) error {
	if nextpc >= memoryCapacity {
		return invalidPCAssignmentError{nextpc}
	}

	if nextpc%2 != 0 {
		return invalidPCAssignmentError{nextpc}
	}

	mem.nextpc = nextpc

	return nil
}

func (mem *Memory) IncPC() {
	mem.pc = mem.nextpc

	if mem.pc >= memoryCapacity {
		mem.pc = 0
	}

	mem.nextpc = mem.pc + ops.InstBytes
}

func (mem *Memory) GetPC() uint16 {
	return mem.pc
}
