package main

// Create a function that reverses a string:
// `test` should be: `tset`
func Reverse(str string) string {
	var reversed = ""
	totalItems := len(str) - 1

	for i := totalItems; i >= 0; i-- {
		reversed = reversed + string(str[i])
	}

	return reversed
}

func main() {
	print(Reverse("test"), "\n")
}
