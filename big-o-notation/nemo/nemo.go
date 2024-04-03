package main

var nemo = []string{"nigel", "dory", "bruce", "marlin", "gill", "bloat", "nemo", "squirt", "darla", "hank"}

func findNemo(array []string) {
	for i, v := range array {
		if v == "nemo" {
			println("Found Nemo at", i)
		}
	}
}

func main() {
	findNemo(nemo) // O(n) --> Linear Time
}
