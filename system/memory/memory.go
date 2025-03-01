package memory

// TODO: Have memory loading/accessing functions for the LDV and VLD instructions, which
// read and write a certain number of registers to and from memory.

const (
	memoryCapacity = 4_096 // 4KB of memory.
	romStart       = 512   // Starting address of Chip-8 programs.
)

type Memory struct {
	ram    [memoryCapacity]byte
	pc     uint16
	nextpc uint16
}

func (mem *Memory) QueueNextPC(nextpc uint16) error {
	if nextpc >= memoryCapacity {
		return outOfBoundsError{nextpc}
	}

	if nextpc%2 != 0 {
		return invalidPCAssignment{nextpc}
	}

	mem.nextpc = nextpc

	return nil
}

func (mem *Memory) IncPC() {
	mem.pc = mem.nextpc

	if mem.pc >= memoryCapacity {
		mem.pc = 0
	}

	mem.nextpc = mem.pc + 2
}

func (mem *Memory) GetPC() uint16 {
	return mem.pc
}
