package main

import "fmt"

func main() {
	strings := []string{"a", "b", "c", "d", "e"}
	fmt.Print(strings[2], "\n")

	// push
	strings = append(strings, "e") // O(1)

	// pop
	strings = strings[0 : len(strings)-2] // O(1)

	//unshift
	strings = append([]string{"x"}, strings...) //O(n)

	//splice
	strings = append(strings[:3], append([]string{"z"}, strings[3:]...)...) //O(n)

	fmt.Print(strings, "\n")
}
