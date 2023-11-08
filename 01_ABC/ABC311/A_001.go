package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	var s string
	fmt.Scan(&n, &s)
	A := strings.IndexByte(s, 'A')
	B := strings.IndexByte(s, 'B')
	C := strings.IndexByte(s, 'C')
	fmt.Println(max(A, B, C) + 1)
}

func max(a ...int) int {
	res := a[0]
	for i := range a {
		if res < a[i] {
			res = a[i]
		}
	}
	return res
}
