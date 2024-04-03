package main

func main() {
	var boxes = []int{1, 2, 3, 4, 5}
	LogPairs(boxes)
}

func LogPairs(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			println(arr[i], arr[j])
		}
	}
}
