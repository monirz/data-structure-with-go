//Singly link list
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
	l.addNode(13)
	l.addNode(14)

	l.addAtFirst("added in first node")

	l.addAfter(2, 12)
	l.traverse()

	fmt.Println(l.size())

	fmt.Println("-----------")
	l.delete(12)
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

func (l *list) addAfter(position int, val interface{}) {
	n := createNode(val)
	right := l.Head
	left := right

	for i := 0; i < position; i++ {
		left = right
		right = right.Next
	}

	left.Next = n
	left = n
	left.Next = right
}

func (l *list) addAtFirst(val interface{}) {
	n := createNode(val)
	n.Next = l.Head

	l.Head = n
}

func (l *list) delete(val interface{}) {
	n := l.Head
	prev := &node{}

	if l.Head == nil {
		fmt.Println("list is empty")
		return
	}

	for {
		if n.Value == val {
			if prev == nil {
				l.Head = n.Next
			} else {
				prev.Next = n.Next
			}

			break
		}

		prev = n
		n = n.Next

		if n.Next == nil {
			fmt.Println("item not found")
			return
		}
	}
}

func (l *list) size() int {

	if l.Head == nil {
		fmt.Println("list is empty")
		return 0
	}

	count := 0
	n := l.Head

	for {
		n = n.Next
		count++
		if n == nil {
			break
		}

	}

	return count
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
