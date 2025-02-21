package main

type OpcodeName int
type Mask = uint16

type Opcode struct {
	Mask Mask
	Code uint16
}

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
// Reg___ opcodes operate between Vx and Vy registers.
// LD_    opcodes load *to* the variable.
// _LD    opcodes load *from* the variable.
const (
	CLS     OpcodeName = iota // 00E0 (Exact)
	RET                       // 00EE (Exact)
	JP                        // 1nnn (Addr)
	CALL                      // 2nnn (Addr)
	SE                        // 3xkk (RegByte)
	SNE                       // 4xkk (RegByte)
	RegSE                     // 5xy0 (TwoReg)
	LD                        // 6xkk (RegByte)
	ADD                       // 7xkk (RegByte)
	RegLD                     // 8xy0 (TwoReg)
	RegOR                     // 8xy1 (TwoReg)
	RegAND                    // 8xy2 (TwoReg)
	RegXOR                    // 8xy3 (TwoReg)
	RegADD                    // 8xy4 (TwoReg)
	RegSUB                    // 8xy5 (TwoReg)
	RegSHR                    // 8xy6 (TwoReg)
	RegSUBN                   // 8xy7 (TwoReg)
	RegSHL                    // 8xyE (TwoReg)
	RegSNE                    // 9xy0 (TwoReg)
	LDI                       // Annn (Addr)
	JPV                       // Bnnn (Addr)
	RND                       // Cxkk (RegByte)
	DRW                       // Dxyn (TwoRegNib)
	SKP                       // Ex9E (Reg)
	SKPNP                     // ExA1 (Reg)
	LDDT                      // Fx07 (Reg)
	LDK                       // Fx0A (Reg)
	DTLD                      // Fx15 (Reg)
	STLD                      // Fx18 (Reg)
	ADDI                      // Fx1E (Reg)
	LDF                       // Fx29 (Reg)
	LDB                       // Fx33 (Reg)
	LDV                       // Fx55 (Reg)
	VLD                       // Fx65 (Reg)
)

func GetOpcodeRef() map[OpcodeName]Opcode {
	return map[OpcodeName]Opcode{
		CLS:     {Exact, 0x00E0},
		RET:     {Exact, 0x00EE},
		JP:      {Addr, 0x1000},
		CALL:    {Addr, 0x2000},
		SE:      {RegByte, 0x3000},
		SNE:     {RegByte, 0x4000},
		RegSE:   {TwoReg, 0x5000},
		LD:      {RegByte, 0x6000},
		ADD:     {RegByte, 0x7000},
		RegLD:   {TwoReg, 0x8000},
		RegOR:   {TwoReg, 0x8001},
		RegAND:  {TwoReg, 0x8002},
		RegXOR:  {TwoReg, 0x8003},
		RegADD:  {TwoReg, 0x8004},
		RegSUB:  {TwoReg, 0x8005},
		RegSHR:  {TwoReg, 0x8006},
		RegSUBN: {TwoReg, 0x8007},
		RegSHL:  {TwoReg, 0x800E},
		RegSNE:  {TwoReg, 0x9000},
		LDI:     {Addr, 0xA000},
		JPV:     {Addr, 0xB000},
		RND:     {RegByte, 0xC000},
		DRW:     {TwoRegNib, 0xD000},
		SKP:     {Reg, 0xE09E},
		SKPNP:   {Reg, 0xE0A1},
		LDDT:    {Reg, 0xF007},
		LDK:     {Reg, 0xF00A},
		DTLD:    {Reg, 0xF015},
		STLD:    {Reg, 0xF018},
		ADDI:    {Reg, 0xF01E},
		LDF:     {Reg, 0xF029},
		LDB:     {Reg, 0xF033},
		LDV:     {Reg, 0xF055},
		VLD:     {Reg, 0xF065},
	}
}
