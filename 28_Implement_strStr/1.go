package main

import (
	"fmt"
	"strings"
)

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	idx := -1
	for i := 0; i < len(haystack); i++ {
		for j := 0; j < len(needle); j++ {
			if i+j >= len(haystack) {
				return -1
			}

			if haystack[i+j] != needle[j] {
				idx = -1
				break
			}
			if j == 0 && haystack[i+j] == needle[j] {
				idx = i
			}
			if j == len(needle)-1 && haystack[i+j] == needle[j] {
				return idx
			}
		}
	}
	return idx
}

func main() {
	haystack := "aaa"
	needle := "aaaa"

	res := strStr2(haystack, needle)
	fmt.Println("res:", res)
}

func strStr2(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}
