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
		length: 1,
	}

	for i := len(list) - 2; i >= 0; i-- {
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
	l.length++
}

func (l *LinkedList[T]) Prepend(item T) {
	node := &Node[T]{
		value: item,
		next:  l.head,
	}

	l.head = node
	l.length++
}

func (l *LinkedList[T]) Insert(index int, item T) {
	var previousNode *Node[T]

	newNode := &Node[T]{
		value: item,
		next:  nil,
	}

	for i := 0; i <= l.length-1 && i <= index-1; i++ {
		if i == 0 {
			previousNode = l.head
		} else {
			previousNode = previousNode.next
		}
	}

	l.length++

	if previousNode == nil {
		newNode.next = l.head
		l.head = newNode
		return
	}

	if previousNode == l.tail {
		l.tail.next = newNode
		newNode.next = nil
		l.tail = newNode
		return
	}

	newNode.next = previousNode.next
	previousNode.next = newNode
}

func (l *LinkedList[T]) ToSlice() []T {
	currentNode := l.head
	slice := []T{}
	for i := 0; currentNode != nil; i++ {
		slice = append(slice, currentNode.value)
		currentNode = currentNode.next
	}

	return slice
}

func main() {
	list := []any{"B", "C", "D"}
	linkedList := NewLinkedList(list)
	linkedList.Prepend("A")
	linkedList.Append("F")
	linkedList.Insert(7, "E")
	// fmt.Print(linkedList.head.next.next.value)
	fmt.Print(linkedList.ToSlice(), " ", linkedList.length, "\n")
}
