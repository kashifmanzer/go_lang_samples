package main

import "fmt"

func main() {
	stack := []int{}

	for i := 0; i < 10; i++ {
		stack = append(stack, i)
	}
	fmt.Println(stack)
	for i := 0; i < 10; i++ {
		if stack[i]%2 == 0 {
			fmt.Printf("%d is even \n", stack[i])
		}
		if stack[i]%2 != 0 {
			fmt.Printf("%d is odd \n", stack[i])
		}
	}
}
