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

// func (inst Instruction) MatchesOpcode(code Opcode) bool {
// 	return uint16(inst)&code.Mask == code.Code
// }

func (inst Instruction) ApplyOpcodeMask(mask Mask) uint16 {
	return uint16(inst) & mask
}
