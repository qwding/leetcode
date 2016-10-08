package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := "111"
	b := "1"
	result := addBinary(a, b)
	fmt.Println("result:", result)
}

//这个方法会导致int 长度不够，肯定不推荐了
func addBinary(a string, b string) string {
	ai, err := strconv.Atoi(a)
	if err != nil {
		return ""
	}

	bi, err := strconv.Atoi(b)
	if err != nil {
		return ""
	}

	result := ""

	c := ai + bi
	fmt.Println("debug: c:", c)

	bit := 0
	for {
		mod := c % 10
		c = c / 10

		if mod > 3 {
			return ""
		} else {
			m := (mod + bit) % 2
			bit = (mod + bit) / 2
			result = strconv.Itoa(m) + result
		}
		if c == 0 {
			if bit == 1 {
				result = "1" + result
			}
			break
		}
	}
	return result
}
