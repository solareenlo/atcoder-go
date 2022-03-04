package main

import (
	"bufio"
	"fmt"
	"os"
)

func d(a, b int) int {
	if a%3 == b%3 {
		return max(a, b)
	}
	return int(1e9)
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	for i := 0; i < n; i++ {
		var r, g, b int
		fmt.Fscan(in, &r, &g, &b)
		r = min(d(g, b), d(b, r), d(r, g))
		if r == int(1e9) {
			fmt.Println(-1)
		} else {
			fmt.Println(r)
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a ...int) int {
	res := a[0]
	for i := range a {
		if res > a[i] {
			res = a[i]
		}
	}
	return res
}
