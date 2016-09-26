package main

import (
	"fmt"
)

func main() {
	arr := "-2147483648"
	result := Atoi(arr)
	fmt.Println("result:", result)
}

func Atoi(str string) int {
	var result int64 = 0
	var sign int64 = 1
	header := false
	for _, s := range []byte(str) {
		if s == ' ' && !header {
			continue
		} else if !header && s == '-' {
			sign = -1
			header = true
			continue
		} else if !header && s == '+' {
			header = true
			sign = 1
			continue
		}
		header = true
		this := s - '0'
		if this >= 10 || this < 0 {
			return int(sign * result)
		}
		result = 10*result + int64(this)
		if result*sign > 2147483647 {
			return 2147483647
		} else if result*sign < -2147483648 {
			return -2147483648
		}
	}
	return int(sign * result)
}
