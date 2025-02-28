package memory

// TODO: Make Program Counter belong to the memory package; it only interfaces
// with accessing memory, therefore it should belong to memory.

// TODO: Implement a 'next program counter' variable that keeps track of the
// next program counter value that will be set. This is needed so that instructions
// that directly set the program counter don't need to offset it by -2 to counteract
// the program counter going up by 2.
// NOTE: This *should* result in the deletion of DecProgramCounter().

// TODO: Have memory loading/accessing functions for the LDV and VLD instructions, which
// read and write a certain number of registers to and from memory.

const (
	memoryCapacity = 4_096                     // 4KB of memory.
	romStart       = 512                       // Starting address of Chip-8 programs.
	romCapacity    = memoryCapacity - romStart // How large a program can be.
)

type Memory struct {
	ram [memoryCapacity]byte
	PC  uint16
}

func (mem *Memory) IncProgramCounter() {
	mem.PC += 2

	if mem.PC >= memoryCapacity-1 {
		mem.PC -= memoryCapacity - 1
	}
}

func (mem *Memory) DecProgramCounter() {
	if mem.PC == 0 {
		mem.PC = memoryCapacity
	}

	mem.PC -= 2
}
