package main

import (
	"bufio"
	"fmt"
	"os"
)

const maxn = 2000 + 10
const D = 1800
const maxD = D*2 + 10

var (
	d = [maxn]int{}
	g = [maxn][maxD]int{}
	n int
)

func no() {
	fmt.Println("No")
	os.Exit(0)
}

func solve(x int, A []int) {
	for i := 1; i <= n; i++ {
		x += d[i]
	}
	if x&1 != 0 {
		no()
	}
	x /= 2
	if x < 0 {
		no()
	}
	breakItem := 0
	totWeight := 0
	for breakItem = 1; breakItem <= n; breakItem++ {
		totWeight += d[breakItem]
		if totWeight > x {
			totWeight -= d[breakItem]
			break
		}
	}
	g := [maxn][maxD]int{}
	for i := range g[breakItem-1] {
		g[breakItem-1][i] = breakItem
	}
	path := [maxn][maxD]int{}
	if D+totWeight-x < 0 {
		no()
	}
	g[breakItem-1][D+totWeight-x] = 0
	for i := breakItem; i <= n; i++ {
		for j := 1; j <= D+D; j++ {
			g[i][j] = g[i-1][j]
		}
		for j := 1; j <= D; j++ {
			if g[i-1][j] < g[i-1][j+d[i]] {
				g[i][j+d[i]] = g[i-1][j]
				path[i][j+d[i]] = i
			}
		}
		for j := D + D; j > D; j-- {
			for k := g[i][j] + 1; k <= g[i-1][j]; k++ {
				if k < g[i][j-d[k]] {
					g[i][j-d[k]] = k
					path[i][j-d[k]] = k
				}
			}
		}
	}
	if g[n][D] == breakItem {
		no()
	}
	tot := D
	for i := 1; i <= n; i++ {
		if i < breakItem {
			A[i] = 1
		} else {
			A[i] = 0
		}
	}
	for i := n; i >= breakItem; i-- {
		for path[i][tot] != 0 {
			cur := path[i][tot]
			A[cur] ^= 1
			if cur < breakItem {
				tot += d[cur]
			} else {
				tot -= d[cur]
				break
			}
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var a, b int
	fmt.Fscan(in, &n, &a, &b)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &d[i])
	}

	x, y := 0, 0
	x += a
	y += a
	x -= b
	y += b
	A := make([]int, maxn)
	B := make([]int, maxn)
	solve(x, A)
	solve(y, B)

	fmt.Fprintln(out, "Yes")
	for i := 1; i <= n; i++ {
		if A[i] != 0 && B[i] != 0 {
			fmt.Fprint(out, "R")
		} else if A[i] != 0 {
			fmt.Fprint(out, "D")
		} else if B[i] != 0 {
			fmt.Fprint(out, "U")
		} else {
			fmt.Fprint(out, "L")
		}
	}
}
