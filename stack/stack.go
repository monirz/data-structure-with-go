//Last In First Out(LIFO)/ stack implementation using array
package main

import (
	"fmt"
)

const size = 10

var stack [size]interface{}
var top = -1

func main() {
}

func push(item interface{}) {
	if isFull() {
		fmt.Println("stack-overflow! stack is full. ")
		return
	}

	top = top + 1
	fmt.Println(top)
	stack[top] = item

}

func pop() interface{} {
	if isEmpty() {
		fmt.Println("stack is empty!")
		return 0
	}

	data := stack[top]
	top = top - 1

	return data
}

func isFull() bool {
	if top == size-1 {
		return true
	}

	return false
}

func isEmpty() bool {
	if top == -1 {
		return true
	}

	return false
}
