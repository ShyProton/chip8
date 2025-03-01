package memory

import (
	"errors"
	"fmt"
)

const stackSize = 16

type callStack struct {
	addresses [stackSize]uint16
	pointer   byte
}

var errStackOverflow = fmt.Errorf("exceeded %d nested calls", stackSize)
var errStackUnderflow = errors.New("cannot pop from empty stack")

func (mem *Memory) PushCallStack() error {
	if mem.cstack.pointer >= stackSize {
		return errStackOverflow
	}

	// The address we return to should be the one after the CALL instruction.
	addr := mem.pc + 2
	if addr >= memoryCapacity {
		addr = 0
	}

	mem.cstack.addresses[mem.cstack.pointer] = addr
	mem.cstack.pointer++

	return nil
}

func (mem *Memory) PopCallStack() error {
	if mem.cstack.pointer == 0 {
		return errStackUnderflow
	}

	mem.cstack.pointer--
	addr := mem.cstack.addresses[mem.cstack.pointer]

	err := mem.QueueNextPC(addr)

	return err
}
