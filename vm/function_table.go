package vm

import "errors"

// FunctionTable maps function names to indices
type FunctionTable struct {
	table map[string]int
}

// NewFunctionTable creates a new function lookup table
func NewFunctionTable() *FunctionTable {
	return &FunctionTable{table: make(map[string]int)}
}

// AddFunction adds a function to the lookup table
func (ft *FunctionTable) AddFunction(name string, index int) {
	ft.table[name] = index
}

// GetFunction retrieves a function index by name
func (ft *FunctionTable) GetFunction(name string) (int, error) {
	index, exists := ft.table[name]
	if !exists {
		return -1, errors.New("function not found")
	}
	return index, nil
}
