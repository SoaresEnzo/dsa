package main

func logFirstTwoBoxes(boxes []int64) {
	println(boxes[0]) // O(1)
	println(boxes[1]) // O(1)
}

func main() {
	boxes := []int64{0, 1, 2, 3, 4, 5}
	logFirstTwoBoxes(boxes) // O(2) --> Constant Time
}
