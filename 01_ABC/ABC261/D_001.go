package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)

	const N = 5005
	x := make([]int, N)
	for i := 1; i <= n; i = i + 1 {
		fmt.Fscan(in, &x[i])
	}

	a := make([]int, N)
	for i := 1; i <= m; i = i + 1 {
		var c, y int
		fmt.Fscan(in, &c, &y)
		a[c] = y
	}

	var f [N][N]int
	b := 0
	for i := 1; i <= n; i = i + 1 {
		f[i][0] = a[0] + b
		b = 0
		for j := 1; j <= i; j = j + 1 {
			f[i][j] = f[i-1][j-1] + x[i] + a[j]
		}
		for j := 0; j <= i; j = j + 1 {
			b = max(b, f[i][j])
		}
	}
	fmt.Println(b)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
