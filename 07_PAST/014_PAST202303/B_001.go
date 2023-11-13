package main

import (
	"fmt"
	"strconv"
)

func main() {
	var d int
	fmt.Scan(&d)
	var a, b int
	var s, t string
	fmt.Scanf("%d.%s", &a, &s)
	fmt.Scanf("%d.%s", &b, &t)
	s = reverseString(s)
	t = reverseString(t)
	a += b
	v := make([]int, d+1)
	for i := 0; i < d; i++ {
		S, _ := strconv.Atoi(string(s[i]))
		T, _ := strconv.Atoi(string(t[i]))
		v[i] += S + T
		v[i+1] += v[i] / 10
		v[i] %= 10
	}

	a += v[d]
	fmt.Printf("%d.", a)
	for i := 0; i < d; i++ {
		fmt.Print(v[d-1-i])
	}
	fmt.Println()
}

func reverseString(s string) string {
	res := []rune(s)
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}
