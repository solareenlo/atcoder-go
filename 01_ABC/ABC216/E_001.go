package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscan(in, &n, &k)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	l, r := -1, int(2e9)
	for l+1 < r {
		m := (l + r) / 2
		x := 0
		for _, y := range a {
			x += max(y-m, 0)
		}
		if x > k {
			l = m
		} else {
			r = m
		}
	}

	x, z := 0, 0
	for _, y := range a {
		x += max(y-r, 0)
		if y > r {
			z += (y*(y+1) - r*(r+1)) / 2
		}
	}
	fmt.Println(z + (k-x)*r)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
