package main

import (
	"bufio"
	"fmt"
	"os"
)

var t int
var A [100001][]int
var B [100001]int

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n, &t)
	for i := 0; i < n-1; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		A[a] = append(A[a], b)
		A[b] = append(A[b], a)
	}
	f(t, -1)
	for i := 1; i <= n; i++ {
		fmt.Printf("%d ", B[i])
	}
}

func f(p, z int) int {
	if len(A[p]) == 1 && p != t {
		return 0
	}
	c := 0
	for _, a := range A[p] {
		if a != z {
			c = max(c, f(a, p))
		}
	}
	B[p] = c + 1
	return c + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
