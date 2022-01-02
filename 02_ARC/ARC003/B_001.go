package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	s := make([]string, n)
	for i := range s {
		fmt.Scan(&s[i])
	}

	for i := range s {
		s[i] = reverseString(s[i])
	}
	sort.Strings(s)
	for i := range s {
		s[i] = reverseString(s[i])
	}

	for i := range s {
		fmt.Println(s[i])
	}
}

func reverseString(s string) string {
	res := []rune(s)
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}
