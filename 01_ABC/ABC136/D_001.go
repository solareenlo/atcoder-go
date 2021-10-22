package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)
	n := len(s)
	res := make([]int, n)

	for j := 0; j < 2; j++ {
		cnt := 0
		for i := 0; i < n; i++ {
			if s[i] == 'R' {
				cnt++
			} else {
				res[i] += cnt / 2
				res[i-1] += (cnt + 1) / 2
				cnt = 0
			}
		}
		res = reverseOrderInt(res)
		s = reverseString(s)
		for i := 0; i < n; i++ {
			if s[i] == 'R' {
				s = s[:i] + "L" + s[i+1:]
			} else {
				s = s[:i] + "R" + s[i+1:]
			}
		}
	}

	for i := 0; i < n; i++ {
		fmt.Print(res[i])
		if i != n-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}

func reverseOrderInt(a []int) []int {
	n := len(a)
	res := make([]int, n)
	n = copy(res, a)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

func reverseString(s string) string {
	res := []rune(s)
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}
