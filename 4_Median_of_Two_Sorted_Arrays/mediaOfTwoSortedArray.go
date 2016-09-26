package main

import (
	"fmt"
)

func main() {
	arr1 := []int{1, 2, 5, 6, 8, 9}
	arr2 := []int{3, 4, 7, 10}
	ret := findMedianSortedArrays(arr1, arr2)
	fmt.Println("result", ret)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	arr := make([]int, len(nums1)+len(nums2))
	for i, j := 0, 0; ; {

	}
}
