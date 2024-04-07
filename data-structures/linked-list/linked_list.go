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

func (l *LinkedList[T]) Print() {
	fmt.Print(
		"Data: ", l.ToSlice(),
		" - Size: ", l.length,
		" - Head: ", l.head.value,
		" - Tail: ", l.tail.value,
		"\n")
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

func (l *LinkedList[T]) getNode(index int) *Node[T] {
	var node *Node[T]
	for i := 0; i <= l.length-1 && i <= index; i++ {
		if i == 0 {
			node = l.head
		} else {
			node = node.next
		}
	}
	return node
}

func (l *LinkedList[T]) Insert(index int, item T) {
	newNode := &Node[T]{
		value: item,
		next:  nil,
	}

	previousNode := l.getNode(index - 1)

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

func (l *LinkedList[T]) Remove(index int) {
	previousNode := l.getNode(index - 1)

	if previousNode == nil {
		l.head = l.head.next
		l.length--
		return
	}

	nodeToBeRemoved := previousNode.next

	if previousNode.next == nil {
		return
	}

	if nodeToBeRemoved == nil || nodeToBeRemoved.next == nil {
		l.tail = previousNode
		previousNode.next = nil
		l.length--
		return
	}

	previousNode.next = previousNode.next.next
	l.length--
	return
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
	linkedList.Print()
	linkedList.Remove(0)
	linkedList.Print()
	linkedList.Remove(4)
	linkedList.Print()

}
