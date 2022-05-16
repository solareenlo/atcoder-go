package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	x := make([]int, 55)
	y := make([]int, 55)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &x[i], &y[i])
	}
	x[n+1] = x[1]
	y[n+1] = y[1]

	mn := make([]int, 55)
	for i := range mn {
		mn[i] = 1 << 60
	}

	kx := make([]int, 55)
	ky := make([]int, 55)
	for i := 1; i <= n; i++ {
		kx[i] = y[i+1] - y[i]
		ky[i] = x[i+1] - x[i]
	}

	var m int
	fmt.Fscan(in, &m)
	for i := 1; i <= m; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		for j := 1; j <= n; j++ {
			mn[j] = min(mn[j], (x[j]+u)*kx[j]-(y[j]+v)*ky[j])
		}
	}

	var Q int
	fmt.Fscan(in, &Q)
	for i := 0; i < Q; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		flag := true
		for i := 1; i <= n; i++ {
			if a*kx[i]-b*ky[i] > mn[i] {
				flag = false
			}
		}
		if flag {
			fmt.Fprintln(out, "Yes")
		} else {
			fmt.Fprintln(out, "No")
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
