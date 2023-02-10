package main

import (
	"bufio"
	"fmt"
	"os"
)

var out = bufio.NewWriter(os.Stdout)

var a [300005]int

func main() {
	defer out.Flush()

	var N, K int
	fmt.Scan(&N, &K)

	if N < K+K {
		fmt.Println(-1)
		return
	}
	for i := 1; i <= N; i++ {
		a[i] = i
	}
	dfs(N, K, 0)
}

func dfs(x, y, z int) {
	if x <= 3*y {
		for i := y + 1; i <= x; i++ {
			fmt.Fprintf(out, "%d ", a[i+z])
		}
		for i := 1; i <= y; i++ {
			fmt.Fprintf(out, "%d ", a[i+z])
		}
		return
	}
	for i := y + 1; i <= y+y; i++ {
		fmt.Fprintf(out, "%d ", a[i+z])
	}
	t := min(y, x-3*y)
	for i := 1; i <= t; i++ {
		fmt.Fprintf(out, "%d ", a[i+z])
	}
	for i := y + y; i >= y+t+1; i-- {
		a[i+z] = a[i-y+z]
	}
	dfs(x-y-t, y, z+y+t)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
