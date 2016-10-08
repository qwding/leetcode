package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := "1010"
	b := "1011"
	result := addBinary(a, b)
	fmt.Println("result:", result)
}

//这个方法会导致int 长度不够，肯定不推荐了
func addBinary(a string, b string) string {
	la := len(a)
	lb := len(b)
	ta := byte('0')
	tb := byte('0')

	result := ""
	i := 0
	bit := byte(0)
	qa := false
	qb := false
	for {
		if la-i-1 >= 0 {
			ta = a[la-i-1]
		} else {
			qa = true
			ta = '0'
		}
		if lb-i-1 >= 0 {
			tb = b[lb-i-1]
		} else {
			qb = true
			tb = '0'
		}

		if qa && qb {
			break
		}

		tmp := ta + tb + bit - 2*'0'
		fmt.Println("ta:", ta, "tb:", tb, "tmp", tmp, "bit:", bit)
		if tmp > 3 {
			return ""
		}
		bit = tmp / 2
		tmp = tmp % 2
		result = strconv.Itoa(int(tmp)) + result
		fmt.Println("result:", result)
		i++
	}
	if bit == byte(1) {
		result = "1" + result
	}
	return result
}
