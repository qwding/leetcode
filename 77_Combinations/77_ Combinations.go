package main

import (
	"fmt"
)

func combine(n int, k int) [][]int {
	// length := getLen(n, k)
	result := make([][]int, 0)
	// addr := &result
	// fmt.Println("len:", len(result), "cap:", cap(result), "addr:", unsafe.Pointer(addr))
	recursive([]int{}, 1, n, k, &result)
	return result
}

func recursive(now []int, num int, n, k int, result *[][]int) {
	for ; num <= n; num++ {
		// addr := &result
		// fmt.Println("debug, num:", num, "now:", now, "result:", result, "addr:", unsafe.Pointer(addr))
		arr := copyArr(now)
		arr = append(arr, num)
		if len(now) == k-1 {
			*result = append(*result, arr)
		} else {
			recursive(arr, num+1, n, k, result)
		}
	}
}

func copyArr(arr []int) []int {
	result := make([]int, len(arr))
	for i, v := range arr {
		result[i] = v
	}
	return result
}

func getLen(n, k int) int {
	if n == k {
		return 1
	}
	a := n
	result := 1
	for ; n > k; n-- {
		result = result * n
		fmt.Println("n:", n, "result:", result)
	}
	result = result / (a - k)
	fmt.Println("n:", n, "k:", k, "result:", result)
	return result
}

func main() {
	n := 13
	k := 10
	fmt.Println("start")
	arr := combine(n, k)
	fmt.Println("arr:", arr)
}
