package memory

func (mem *Memory) GetInstBytes() (byte, byte) {
	return mem.ram[mem.pc], mem.ram[mem.pc+1]
}

func (mem *Memory) ByteAt(i uint16) (*byte, error) {
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
