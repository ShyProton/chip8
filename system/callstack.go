package system

import (
	"errors"
	"fmt"
)

const stackSize = 16

type CallStack struct {
	stack [stackSize]uint16
	sp    byte
}

var errStackOverflow = fmt.Errorf("exceeded %d nested calls", stackSize)
var errStackUnderflow = errors.New("cannot pop from empty stack")

func (callstack *CallStack) Push(pc uint16) error {
	if callstack.sp >= stackSize {
		return errStackOverflow
	}

	callstack.stack[callstack.sp] = pc
	callstack.sp++

	return nil
}

func (callstack *CallStack) Pop() (uint16, error) {
	if callstack.sp <= 0 {
		return 0, errStackUnderflow
	}

	callstack.sp--

	return callstack.stack[callstack.sp], nil
}
