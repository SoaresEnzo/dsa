package main

// Given 2 arrays, create a gunction that let's a user know (true/false) whether these two arrays contains any common items
// For Example:
// array1 = [a,b,c,x]
// array2 = [z,y,i]
// should return false
//------------------
// array1 = [a,b,c,x]
// array2 = [z,y,x]
// should return true

func main() {
	array1 := []string{"a", "b", "c", "x"}
	array2 := []string{"z", "y", "c"}

	print(containsCommonItem(array1, array2), "\n")
}

// func containsCommonItem(array1 []string, array2 []string) bool {
// 	for _, fValue := range array1 {
// 		for _, sValue := range array2 {
// 			if fValue == sValue {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }

func containsCommonItem(array1 []string, array2 []string) bool {
	items := map[string]bool{}
	for _, value := range array1 {
		items[value] = true
	}

	for _, value := range array2 {
		if items[value] {
			return true
		}
	}

	return false
}
