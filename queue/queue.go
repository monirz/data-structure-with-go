//Array based First In First Out(FIFO)/Queue
package main

import "fmt"

const MAX = 5

var queue [MAX]interface{}
var rear = -1
var front = -1
var size int

func main() {

}

func enQueue(item interface{}) error {
	if isFull() {
		fmt.Println("queue is full!")
		return fmt.Errorf("queue is full")
	}

	if front == -1 {
		front = 0
	}

	rear = (rear + 1) % MAX
	size++
	queue[rear] = item

	return nil

}

func deQueue() interface{} {
	if isEmpty() {
		fmt.Println("queue is empty !")
		return 0
	}

	data := queue[front]

	if front == rear {
		front = -1
		rear = -1
	} else {
		front = (front + 1) % MAX
		size--
	}

	return data

}

func isFull() bool {
	if front == 0 && rear == MAX-1 {
		return true
	}

	return false
}

func isEmpty() bool {
	if front == -1 {
		return true
	}

	return false
}
