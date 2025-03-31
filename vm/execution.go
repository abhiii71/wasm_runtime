package vm

import (
	"errors"
	"fmt"

	"github.com/abhiii71/wasm_runtime/parser"
)

// ExecuteWASMFunction runs a parsed  WASM function
func ExecuteWASMFunction(stack *Stack, function parser.WASMFunction) error {
	for i := 0; i < len(function.Code); i++ {
		opcode := function.Code[i]

		switch opcode {
		case 0x41: // i32.const
			i++
			value := int64(function.Code[i])
			stack.Push(value)
			fmt.Println("Pushed: ", value)
		case 0x6A: // i32.add
			val1, err1 := stack.Pop()
			val2, err2 := stack.Pop()
			if err1 != nil || err2 != nil {
				return errors.New("stack underflow")
			}
			stack.Push(val1 + val2)
			fmt.Println("Added:", val1, "+", val2, "=", val1+val2)

		case 0x6B: // i32.sub
			val1, err1 := stack.Pop()
			val2, err2 := stack.Pop()
			if err1 != nil || err2 != nil {
				return errors.New("stack underflow")
			}
			stack.Push(val1 - val2)
			fmt.Println("Subtracted: ", val1, "-", val2, "=", val1-val2)
		case 0x0B: // end
			fmt.Println("Function execution compelete.")
			return nil

		default:
			return fmt.Errorf("unsupported opcode: 0x%x", opcode)

		}
	}
	return nil
}
