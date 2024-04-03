package main

// Find 1st, Find Nth

func main() {
	tweets := []string{"hi", "my", "teddy"}
	println(tweets[0])             // O(1)
	println(tweets[len(tweets)-1]) // O(1)
}
