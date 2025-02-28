package system

type Opcode = uint16
type Mask = uint16

// Possible mask types.
const (
	Exact     Mask = 0xFFFF // ____
	Addr      Mask = 0xF000 // _nnn
	Reg       Mask = 0xF0FF // _x__
	RegByte   Mask = 0xF000 // _xkk
	TwoReg    Mask = 0xF00F // _xy_
	TwoRegNib Mask = 0xF000 // _xyn
)

// Reference: http://devernay.free.fr/hacks/chip8/C8TECH10.HTM#3.1.
//
// The Opcode identifiers, which does not include specific arguments passed to
// the instruction; meaning these are to be matched after applying one of the
// masks onto an instruction.
//
// Reg___ opcodes operate between Vx and Vy registers.
//
// LD_    opcodes load *to* the variable.
//
// _LD    opcodes load *from* the variable.
const (
	CLS     Opcode = 0x00E0 // 00E0 (Exact)
	RET     Opcode = 0x00EE // 00EE (Exact)
	JP      Opcode = 0x1000 // 1nnn (Addr)
	CALL    Opcode = 0x2000 // 2nnn (Addr)
	SE      Opcode = 0x3000 // 3xkk (RegByte)
	SNE     Opcode = 0x4000 // 4xkk (RegByte)
	RegSE   Opcode = 0x5000 // 5xy0 (TwoReg)
	LD      Opcode = 0x6000 // 6xkk (RegByte)
	ADD     Opcode = 0x7000 // 7xkk (RegByte)
	RegLD   Opcode = 0x8000 // 8xy0 (TwoReg)
	RegOR   Opcode = 0x8001 // 8xy1 (TwoReg)
	RegAND  Opcode = 0x8002 // 8xy2 (TwoReg)
	RegXOR  Opcode = 0x8003 // 8xy3 (TwoReg)
	RegADD  Opcode = 0x8004 // 8xy4 (TwoReg)
	RegSUB  Opcode = 0x8005 // 8xy5 (TwoReg)
	RegSHR  Opcode = 0x8006 // 8xy6 (TwoReg)
	RegSUBN Opcode = 0x8007 // 8xy7 (TwoReg)
	RegSHL  Opcode = 0x800E // 8xyE (TwoReg)
	RegSNE  Opcode = 0x9000 // 9xy0 (TwoReg)
	LDI     Opcode = 0xA000 // Annn (Addr)
	JPV     Opcode = 0xB000 // Bnnn (Addr)
	RND     Opcode = 0xC000 // Cxkk (RegByte)
	DRW     Opcode = 0xD000 // Dxyn (TwoRegNib)
	SKP     Opcode = 0xE09E // Ex9E (Reg)
	SKPNP   Opcode = 0xE0A1 // ExA1 (Reg)
	LDDT    Opcode = 0xF007 // Fx07 (Reg)
	LDK     Opcode = 0xF00A // Fx0A (Reg)
	DTLD    Opcode = 0xF015 // Fx15 (Reg)
	STLD    Opcode = 0xF018 // Fx18 (Reg)
	ADDI    Opcode = 0xF01E // Fx1E (Reg)
	LDF     Opcode = 0xF029 // Fx29 (Reg)
	LDB     Opcode = 0xF033 // Fx33 (Reg)
	LDV     Opcode = 0xF055 // Fx55 (Reg)
	VLD     Opcode = 0xF065 // Fx65 (Reg)
)
