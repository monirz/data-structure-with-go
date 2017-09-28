package main

import "fmt"

const size = 5

var queue [size]interface{}
var rear = -1
var front = -1

func main() {
	enQueue(121)
	enQueue(1120)
	enQueue(1222)
	enQueue(11)
	enQueue("Go FTW!")
	deQueue()
	enQueue("another item")

	fmt.Println(queue)

}

func enQueue(item interface{}) {
	if isFull() {
		fmt.Println("queue is full!")
		return
	}

	if front == -1 {
		front = 0
	}

	rear = (rear + 1) % size
	queue[rear] = item

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
		front = (front + 1) % size
	}

	return data

}

func isFull() bool {
	if front == 0 && rear == size-1 {
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
