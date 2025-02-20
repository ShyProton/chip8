package main

import (
	"errors"
	"fmt"
)

const StackSize = 16

type CallStack [StackSize]uint16

var ErrStackOverflow = fmt.Errorf("exceeded %d nested calls", StackSize)
var ErrStackUnderflow = errors.New("cannot pop from empty stack")

func (stack *CallStack) Push(reg *Registers) error {
	if reg.SP >= StackSize {
		return ErrStackOverflow
	}

	stack[reg.SP] = reg.PC
	reg.SP++

	return nil
}

func (stack *CallStack) Pop(reg *Registers) error {
	if reg.SP <= 0 {
		return ErrStackUnderflow
	}

	reg.SP--
	reg.PC = stack[reg.SP]

	return nil
}
