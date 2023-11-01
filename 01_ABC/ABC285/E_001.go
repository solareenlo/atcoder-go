package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var x, y, z [5050]int

	var n int
	fmt.Fscan(in, &n)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &x[i])
		x[i] += x[i-1]
		y[i] = x[(i+1)>>1] + x[i>>1]
	}
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			z[i] = max(z[i], z[j]+y[i-j-1])
		}
	}
	fmt.Println(z[n])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
