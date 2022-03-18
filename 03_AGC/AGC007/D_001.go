package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m, t int
	fmt.Fscan(in, &n, &m, &t)

	const N = 100010
	x := make([]int, N)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &x[i])
	}

	k := 1 << 60
	f := make([]int, N)
	for i, j := 1, 0; i <= n; i++ {
		for ; 2*(x[i]-x[j+1]) > t; j++ {
			k = min(k, f[j]-2*x[j+1])
		}
		f[i] = min(k+2*x[i], f[j]+t)
	}

	fmt.Println(f[n] + m)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
