package main

import "fmt"

type node struct {
	Value interface{}
	Next  *node
}

type list struct {
	Head *node
	Tail *node
}

func main() {

	l := &list{}
	l.addNode(11)
	l.addNode(12)
	l.addNode(14)

	l.traverse()

}

func createNode(val interface{}) *node {
	n := &node{}
	n.Value = val

	return n
}

func (l *list) addNode(val interface{}) {

	n := createNode(val)

	if l.Head == nil {
		l.Head = n
		l.Tail = n
	} else {
		l.Tail.Next = n
		l.Tail = n
	}

}

func (l *list) traverse() {

	if l.Head == nil {
		fmt.Println("list is empty")
		return
	}

	n := l.Head

	for {
		fmt.Println(n.Value)
		n = n.Next
		if n == nil {
			break
		}
	}

}
