package parser

import (
	"bytes"
	"encoding/binary"
	"errors"
	"os"
)

// WASMModule represents a parsed WASM module
type WASMModule struct {
	Functions []WASMFunction
}

// WASMFunction represents a function in the WASM module
type WASMFunction struct {
	Code []byte // Raw WASM bytecode
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
	index := 8 // skipping magic & version

	// Find function section (0x03) and code section (0x0A)
	for index < len(data) {
		sectionID := data[index]
		index++
		sectionSize := int(binary.LittleEndian.Uint32(data[index : index+4]))
		index += 4

		if sectionID == 0x0A { // Code section
			functionCount := int(data[index]) // Number of functions
			index++

			for i := 0; i < functionCount; i++ {
				codeSize := int(data[index]) // Function size
				index++
				module.Functions = append(module.Functions, WASMFunction{
					Code: data[index : index+codeSize],
				})
				index += codeSize
			}
		} else {
			index += sectionSize // skip irrelevant sections
		}
	}
	return module, nil
}
