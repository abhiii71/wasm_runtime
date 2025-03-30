package tests

import (
	"testing"

	"github.com/abhiii71/wasm_runtime/vm"
)

func TestStack(t *testing.T) {
	stack := vm.NewStack()
	stack.Push(100)
	val, _ := stack.Pop()
	if val != 100 {
		t.Errorf("Expected 100, got %d", val)
	}
}
