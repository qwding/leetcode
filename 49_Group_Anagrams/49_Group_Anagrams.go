package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	result := [][]string{}
	for i, _ := range strs {
		s := sortStr(strs[i])
		has := false
		for j := 0; j < len(result); j++ {
			if result[j][0] == s {
				has = true
				result[j] = append(result[j], strs[i])
				break
			}
		}
		if !has {
			result = append(result, []string{s, strs[i]})
		}
	}

	for i := 0; i < len(result); i++ {
		result[i] = result[i][1:]
	}

	return result
}

func main() {
	// arr := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	arr := []string{}
	result := groupAnagrams(arr)
	fmt.Println("result:", result)
}

func sortStr(str string) string {
	arr := []int{}

	for _, s := range str {
		arr = append(arr, int(s))
	}
	sort.Ints(arr)
	result := ""
	for _, s := range arr {
		result = result + string(s)
	}
	return result
}
