package tests

import (
	"testing"

	"github.com/abhiii71/wasm_runtime/memory"
	"github.com/abhiii71/wasm_runtime/vm"
)

func TestExecutionEngine(t *testing.T) {
	bytecode := []byte{0x41, 10, 0x41, 20, 0x6A, 0x0B} // i32.const 10, i32.const 20, i32.add, end

	funcTable := vm.NewFunctionTable()
	mem := memory.NewMemory(1024)
	engine := vm.NewExecutionEngine(bytecode, funcTable, mem)

	err := engine.Run()
	if err != nil {
		t.Errorf("Execution failed: %v", err)
	}

	// Use the getter method
	result, err := engine.GetStack().Pop()
	if err != nil || result != 30 {
		t.Errorf("Expected 30, got %d", result)
	}
}
