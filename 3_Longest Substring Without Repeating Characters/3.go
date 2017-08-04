package main

import (
	"fmt"
)

func main() {
	str := "c"         //1
	str2 := "abcabcbb" //3
	str3 := "pwwkew"   //3
	str4 := "bbbbb"    //1
	str5 := "dvdf"     //3
	a := lengthOfLongestSubstring(str)
	b := lengthOfLongestSubstring(str2)
	c := lengthOfLongestSubstring(str3)
	d := lengthOfLongestSubstring(str4)
	e := lengthOfLongestSubstring(str5)

	fmt.Println("res:", a, b, c, d, e)
}

func lengthOfLongestSubstring(s string) int {
	max := 0

	tmplen := 0
	head := 0
	same := -1

	for i := 0; i < len(s); i++ {
		has := false
		for t := head; t < i; t++ {
			if s[t] == s[i] {
				has = true
				same = t
				break
			}
		}
		if !has {
			tmplen++
		} else {
			if max < tmplen {
				max = tmplen
			}
			head = same + 1
			tmplen = i - same
		}
	}

	if max < tmplen {
		max = tmplen
	}

	return max
}
