package main

import (
	"./ds"
	"fmt"
)

func stack_test() {
	stack := new(ds.Stack)
	stack.Push(1)
	stack.Push("abc")
	fmt.Println(stack.Top())
	stack.Pop()
	fmt.Println(stack.Top())
}
func main() {
	stack_test()
}
