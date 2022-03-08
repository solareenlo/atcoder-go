package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	ans int
	x   = [505]int{}
	y   = [505]int{}
	a   = [505][505]int{}
)

func add(i, j, k int) {
	x[i] -= k
	y[j] -= k
	a[i][j] += k
	ans += abs(k)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)
	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			fmt.Fscan(in, &a[i][j])
			if (i+j)&1 != 0 {
				a[i][j] *= -1
			}
			x[i] += a[i][j]
			y[j] += a[i][j]
		}
	}

	for i := range a {
		for j := range a[i] {
			a[i][j] = 0
		}
	}
	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			if x[i] < 0 && y[j] < 0 {
				add(i, j, max(x[i], y[j]))
			}
			if x[i] > 0 && y[j] > 0 {
				add(i, j, min(x[i], y[j]))
			}
		}
	}
	for i := 1; i < n+1; i++ {
		add(i, 1, x[i])
	}
	for i := 1; i < m+1; i++ {
		add(1, i, y[i])
	}

	fmt.Fprintln(out, ans)
	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			if (i+j)&1 != 0 {
				fmt.Fprint(out, -a[i][j], " ")
			} else {
				fmt.Fprint(out, a[i][j], " ")
			}
		}
		fmt.Fprintln(out)
	}
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
