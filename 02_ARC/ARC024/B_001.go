package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(in, &N)

	c := make([]int, N)
	for i := range c {
		fmt.Fscan(in, &c[i])
	}

	res := 0
	n := 1
	for i := 1; i < 2*N; i++ {
		if c[(i-1)%N] == c[i%N] {
			n++
		} else {
			res = max(res, n)
			n = 1
		}
	}
	if res == 0 {
		fmt.Println(-1)
	} else {
		fmt.Println((res + 1) / 2)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
