package main

import (
	"fmt"
)

func permute(nums []int) [][]int {
	length := GetLen(len(nums))

	result := make([][]int, length)
	recurrence(nums, result, 0, []int{})

	return result
}

func GetLen(length int) int {
	result := 1
	for i := 1; i <= length; i++ {
		result = result * i
	}
	return result
}

func recurrence(nums []int, result [][]int, index int, tmp []int) int {
	// fmt.Println("nums:", nums, "index:", index, "tmp:", tmp)
	for i := 0; i < len(nums); i++ {
		tmp2 := append(tmp, nums[i])
		if len(nums) == 1 {
			// fmt.Println("in return")
			result[index] = tmp2
			index++
		} else {

			// slice 的append原理是这样的。 如果cap可以装下，那么结果地址不变，但是如果cap装不下，地址变
			arr := make([]int, len(nums))
			for j := 0; j < len(nums); j++ {
				arr[j] = nums[j]
			}
			// fmt.Println("arr:", arr)

			arr = append(arr[:i], arr[i+1:]...)
			// fmt.Println("i:", i, "arr:", arr, "nums:", nums)
			index = recurrence(arr, result, index, tmp2)
		}
	}
	return index
}

func main() {
	nums := []int{1, 2, 3, 4}

	result := permute(nums)
	fmt.Println("result:", result)

}
