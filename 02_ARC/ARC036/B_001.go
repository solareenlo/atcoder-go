package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	h := make([]int, n)
	for i := range h {
		fmt.Fscan(in, &h[i])
	}

	l := make([]int, n)
	for i := 0; i < n-1; i++ {
		if h[i] < h[i+1] {
			l[i+1]++
		}
	}

	r := make([]int, n)
	for i := n - 1; i > 0; i-- {
		if h[i-1] > h[i] {
			r[i-1]++
		}
	}

	c := make([]int, n)
	for i := 0; i < n; i++ {
		if l[i]+r[i] > 0 {
			c[i]++
		}
	}

	res := 0
	sum := 0
	c[0] = 0
	c[n-1] = 0
	for i := 0; i < n; i++ {
		if c[i] > 0 {
			sum++
			res = max(res, sum)
		} else {
			sum = 0
		}
	}

	if n == 1 {
		fmt.Println(1)
	} else {
		fmt.Println(res + 2)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
