package parser

import (
	"bytes"
	"encoding/binary"
	"errors"
	"os"
)

// WASMModule represents a parsed WASM file
type WASMModule struct {
	Functions []int
}

// ParseWASMFile reads and parses a .wasm file
func ParseWASMFile(filename string) (*WASMModule, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	if !bytes.Equal(data[:4], []byte("\x00asm")) {
		return nil, errors.New("invalid WASM file")
	}

	module := &WASMModule{}
	// Read function section (simplified)
	for i := 8; i < len(data); i++ {
		module.Functions = append(module.Functions, int(binary.LittleEndian.Uint32(data[i:i+4])))
	}
	return module, nil
}
