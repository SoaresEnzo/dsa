package main

import "fmt"

type LinkedList[T any] struct {
	head   *Node[T]
	tail   *Node[T]
	length int
}

type Node[T any] struct {
	value T
	next  *Node[T]
}

func NewLinkedList(list []any) *LinkedList[any] {

	firstNode := &Node[any]{
		value: list[len(list)-1],
		next:  nil,
	}

	linkedList := &LinkedList[any]{
		tail:   firstNode,
		head:   firstNode,
		length: 0,
	}

	for i := len(list) - 1; i >= 0; i-- {
		linkedList.length++
		node := &Node[any]{
			value: list[i],
			next:  linkedList.head,
		}
		linkedList.head = node
	}

	return linkedList
}

func (l *LinkedList[T]) Append(item T) {
	node := &Node[T]{
		value: item,
		next:  nil,
	}

	l.tail.next = node
	l.tail = node
}

func main() {
	list := []any{"A", "B", "C"}
	linkedList := NewLinkedList(list)
	linkedList.Append("D")
	fmt.Print(linkedList.tail.value)
}
