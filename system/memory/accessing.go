package memory

func (mem *Memory) GetInstBytes() (byte, byte) {
	return mem.ram[mem.pc], mem.ram[mem.pc+1]
}

func (mem *Memory) ByteAt(i int) (*byte, error) {
	if i >= memoryCapacity {
		return nil, outOfBoundsError{i}
	}

	return &mem.ram[i], nil
}

func (mem *Memory) FontAddr(digit byte) (uint16, error) {
	if digit >= fontChars {
		return 0, invalidFontAccess{digit}
	}

	return uint16(digit) * fontCharRows, nil
}

func (mem *Memory) ReadToBytes(addr int, bytes []byte) error {
	// Cannot read registers from memory if we exceed the memory capacity while doing so.
	largestAddr := addr + len(bytes) - 1
	if largestAddr >= memoryCapacity {
		return outOfBoundsError{memoryCapacity}
	}

	for i := range bytes {
		bytes[i] = mem.ram[addr+i]
	}

	return nil
}
