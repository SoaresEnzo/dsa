package main

import "fmt"

type Array[T interface{}] struct {
	data   map[int]T
	length int
}

func (a *Array[T]) Get(index int) T {
	return a.data[index]
}

func (a *Array[T]) Push(item T) {
	a.data[a.length] = item
	a.length++
}

func (a *Array[T]) Pop() {
	delete(a.data, a.length-1)
	a.length--
}

func (a *Array[T]) Delete(index int) {
	a.shiftItems(index)
}

func (a *Array[T]) shiftItems(index int) {
	for i := index; i < a.length-1; i++ {
		a.data[i] = a.data[i+1]
	}
	a.Pop()
}

func main() {
	myArray := Array[string]{
		data:   make(map[int]string),
		length: 0,
	}

	myArray.Push("a")
	myArray.Push("hi")
	myArray.Push("World")
	myArray.Delete(1)
	fmt.Print(myArray)
}
