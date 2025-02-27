package main

const (
	InstBits      = 16
	BitsPerByte   = 8
	InstBytes     = InstBits / BitsPerByte
	BitsPerNibble = BitsPerByte / 2
	InstNibbles   = InstBytes * 2
)

type Instruction uint16

func InstFromBytes(even, odd byte) Instruction {
	hiHalf := uint16(even) << uint16(BitsPerByte)
	loHalf := uint16(odd)

	return Instruction(hiHalf | loHalf)
}

func (inst Instruction) GetHexDigits() [InstNibbles]uint {
	const FirstHexMask = 0xF

	var digits [InstNibbles]uint

	for i := range InstNibbles {
		digits[i] = uint(inst) >> (BitsPerNibble * i) & FirstHexMask
	}

	return digits
}

func (inst Instruction) ApplyOpcodeMask(mask Mask) Opcode {
	return uint16(inst) & mask
}

// Gets the address component of an Address instruction.
func (inst Instruction) GetAddr() uint16 {
	// Applying the NOT of the address mask gets the address parameter passed
	// into the instruction.
	return uint16(inst) & ^Addr
}

// Gets the register and byte components of a RegByte instruction.
func (inst Instruction) GetRegByte() (uint16, byte) {
	const (
		RegIdx   = 2      // Need two shifts to move Reg to the right.
		ByteMask = 0x00FF // Last two digits are the byte argument.
	)

	params := uint16(inst) & ^RegByte

	reg := params >> (RegIdx * BitsPerNibble)
	b := byte(params & ByteMask)

	return reg, b
}

func (inst Instruction) GetTwoReg() (uint16, uint16) {
	const (
		RegIdxX       = 2      // Need two shifts to move RegX to the right.
		RegIdxY       = 1      // Need one shift to move RegY to the right.
		LastDigitMask = 0x000F // Need the last digit to isolate RegY from RegX.
	)

	params := uint16(inst) & ^TwoReg

	regx := params >> (RegIdxX * BitsPerNibble)
	regy := params >> (RegIdxY * BitsPerNibble) & LastDigitMask

	return regx, regy
}
