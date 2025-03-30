WebAssembly (WASM) Runtime in Go

🚀 Project Overview

This project is a WebAssembly (WASM) runtime implemented in Go, designed to execute WASM bytecode using a stack-based execution model. The runtime includes parsing of WASM files, function lookups using hash tables, and execution of simple functions.

📂 Project Structure

wasm_runtime/
│── main.go                  # Entry point of the runtime
│── parser/
│   ├── wasm_parser.go       # Parses .wasm files and extracts function details
│── vm/
│   ├── stack.go             # Implements the stack data structure for execution
│   ├── function_table.go    # Implements hash table for function lookup
│   ├── execution.go         # Handles execution of WASM functions
│── memory/
│   ├── memory.go            # Manages linear memory for WASM execution
│── tests/
│   ├── example_test.go      # Unit tests for various components
│── go.mod                   # Go module dependencies
│── go.sum                   # Go module lockfile
│── README.md                # Project documentation

📌 Features Implemented (Day 1)

✅ Stack-based execution using a custom stack implementation.✅ Hash table for function lookups to quickly retrieve function indices.✅ Parsing of WASM files to extract function information.✅ Simple function execution using the stack and lookup table.✅ Unit tests for stack operations.

🛠 Installation & Setup

1️⃣ Clone the repository

git clone https://github.com/your-repo/wasm_runtime.git
cd wasm_runtime

2️⃣ Install dependencies

go mod tidy

3️⃣ Run the tests

go test ./tests

🔧 Components Breakdown

Stack (LIFO) - stack.go

Used for executing WASM instructions.

Operations: Push, Pop, Peek.

Example:

stack := NewStack()
stack.Push(10)
val, _ := stack.Pop() // Returns 10

Function Lookup Table (Hash Table) - function_table.go

Stores function names and their respective indices.

Example:

ft := NewFunctionTable()
ft.AddFunction("add", 1)
index, _ := ft.GetFunction("add") // Returns 1

WASM File Parser - wasm_parser.go

Reads .wasm files and extracts function details.

Example:

module, err := ParseWASMFile("test.wasm")
fmt.Println("Parsed functions:", module.Functions)

Function Execution - execution.go

Executes functions based on the lookup table.

Example:

stack := NewStack()
ft := NewFunctionTable()
ft.AddFunction("add", 1)
ExecuteWASMFunction(stack, ft, "add")

🎯 Future Enhancements

Support for executing more complex WASM instructions.

Implementing linear memory operations.

Optimizing function execution with JIT compilation.

Expanding unit tests for better coverage.

📢 Contributions are welcome! Feel free to open issues and submit PRs.


