package main

import (
	"fmt"
)

func main() {
	arr1 := []int{1, 2, 5, 6, 8, 9}
	arr2 := []int{3, 4, 7, 10}
	ret := findMedianSortedArrays(arr1, arr2)
	fmt.Println("result", ret)

	// fmt.Println("res:", findindx(arr2, 10))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	head := 0
	tail := len(arr1) - 1

	head2:=0
	tail2:= len(arr2) -1

	const num := (tail+tail2)/2
	ji := 0
	if (tail+tail2)%2 == 0{
		ji = 1
	}

	for {
		mid := (head + tail) / 2
		idx := findindx(arr2[head2:tail2], arr1[mid])
		if mid+idx ==  num {
			
		}
	}

}

func findindx(arr []int, target int) int {
	length := len(arr)
	if length == 0 {
		return 0
	}
	head := 0
	tail := length - 1
	count := 0
	for {
		fmt.Println("debug:", head, tail)
		if tail == head {
			return head
		}
		if tail == head+1 {
			if arr[tail] <= target {
				return tail
			} else {
				return head
			}
		}

		if count > 5 {
			return head
		} else {
			count++
		}

		if target == arr[(head+tail)/2] {
			return (head + tail) / 2
		} else if target > arr[(head+tail)/2] {
			head = (head + tail) / 2
		} else {
			tail = (head + tail) / 2
		}
	}
}
