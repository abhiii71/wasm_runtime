package vm

import (
	"errors"
	"fmt"

	"github.com/abhiii71/wasm_runtime/memory"
)

type ExecutionEngine struct {
	stack    *Stack         // stack for execution
	memory   *memory.Memory // Memory for data storage
	funcs    *FunctionTable // Function lookup table
	bytecode []byte         // Load WASM bytecode
	pc       int            // Program counter
}

// NewExecutionEngine intializes a new WASM execution engine
func NewExecutionEngine(bytecode []byte, funcs *FunctionTable, memory *memory.Memory) *ExecutionEngine {
	return &ExecutionEngine{
		stack:    NewStack(),
		memory:   memory,
		funcs:    funcs,
		bytecode: bytecode,
		pc:       0,
	}
}

// GetStack returns the execution stack for testing purposes
func (e *ExecutionEngine) GetStack() *Stack {
	return e.stack
}

// Run executes the WASM bytecode
func (e *ExecutionEngine) Run() error {
	for e.pc < len(e.bytecode) {
		opcode := e.bytecode[e.pc] // Fetch instruction
		e.pc++                     // Move to next  instruction

		switch opcode {
		case 0x41: // i32.const
			value := int(e.bytecode[e.pc])
			e.pc++
			e.stack.Push(value)

		case 0x6A: // i32.add
			b, err := e.stack.Pop()
			if err != nil {
				return errors.New("stack underflow")
			}
			a, err := e.stack.Pop()
			if err != nil {
				return errors.New("stack underflow")
			}
			e.stack.Push(a + b)

		case 0x28: // Load from memory(1.32.load)
			address, err := e.stack.Pop()
			if err != nil {
				return errors.New("stack underflow")
			}
			value, err := e.memory.Load(address)
			if err != nil {
				return err
			}
			e.stack.Push(value)
		case 0x36: // store to memory(i32.store)
			value, err := e.stack.Pop()
			if err != nil {
				return errors.New("stack underflow")
			}
			address, err := e.stack.Pop()
			if err != nil {
				return errors.New("stack underflow")
			}
			err = e.memory.Store(address, value)
			if err != nil {
				return err
			}
		case 0x0B: // end of function
			fmt.Println("Function execution compelete.")
			return nil

		default:
			return fmt.Errorf("unsupported opcode: 0x%x", opcode)

		}
	}
	return nil
}
