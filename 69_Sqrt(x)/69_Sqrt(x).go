package main

import (
	"fmt"
)

func mySqrt(x int) int {
	flag := 1
	if x < 0 {
		flag = -1
		x = -1 * x
	} else if x == 0 {
		return 0
	} else if x >= 1 && x < 4 {
		return 1
	} else if x > -4 && x <= -1 {
		return -1
	}

	r := x / 2
	fmt.Println("r := ", r)
	i := 2
	for ; i <= r; i++ {
		if i*i < x {
			continue
		} else if i*i == x {
			return i
		} else {
			return (i - 1) * flag
		}
	}
	return i - 1
}

func main() {
	fmt.Println("start")
	test := 9
	result := mySqrt(test)
	fmt.Println("result:", result)

}
