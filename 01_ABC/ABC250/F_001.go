package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	x = make([]int, 400008)
	y = make([]int, 400008)
)

func area(i, j, k int) int {
	return (x[j]-x[i])*(y[k]-y[i]) - (x[k]-x[i])*(y[j]-y[i])
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &x[i], &y[i])
		x[i+n] = x[i]
		y[i+n] = y[i]
	}

	S := 0
	A := 0
	ans := 1 << 61
	for i := 1; i <= n; i++ {
		S += x[i]*y[i+1] - x[i+1]*y[i]
	}

	for i, j := 1, 2; i <= n; i++ {
		for A < S {
			A += 4 * area(i, j, j+1)
			ans = min(ans, abs(A-S))
			j++
		}
		A -= 4 * area(j, i, i+1)
		ans = min(ans, abs(A-S))
	}
	fmt.Println(ans)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
