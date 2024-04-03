package main

import (
	"fmt"
	"slices"
)

func MergeSortedArrays(nums1 []int, nums2 []int) []int {
	n := len(nums1)
	value := append(nums1[:n], nums2...)
	slices.Sort(value)
	return value
}

func main() {
	arr1 := []int{}
	arr2 := []int{1}

	fmt.Print(MergeSortedArrays(arr1, arr2))
}
