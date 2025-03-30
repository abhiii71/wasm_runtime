package vm

import "fmt"

func ExecuteWASMFunction(stack *Stack, ft *FunctionTable, funcName string) {
	funcIndex, err := ft.GetFunction(funcName)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println("Executing Function: ", funcName, "at index", funcIndex)

	stack.Push(42)
	result, _ := stack.Pop()
	fmt.Println("Function result:", result)
}
