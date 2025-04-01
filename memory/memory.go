package memory

import (
	"errors"
)

// Memory represents a WASM linear memory
type Memory struct {
	data []int
}

// NewMemory intializes memory with a given size
func NewMemory(size int) *Memory {
	return &Memory{data: make([]int, size)}
}

// Load fetches a value from memory
func (m *Memory) Load(address int) (int, error) {
	if address < 0 || address >= len(m.data) {
		return 0, errors.New("memory out of bounds")
	}
	return m.data[address], nil
}

// Store writes a value to memory
func (m *Memory) Store(address, value int) error {
	if address < 0 || address >= len(m.data) {
		return errors.New("memory out of bounds")
	}
	m.data[address] = value
	return nil
}
