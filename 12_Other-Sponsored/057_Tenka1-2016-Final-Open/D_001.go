package main

import (
	"bufio"
	"fmt"
	"os"
)

const MN = 20

var n, C, D int
var cost [2][MN]int
var g [MN]int
var dp [1 << MN][MN][2]int
var used [1 << MN][MN][2]bool

func main() {
	in := bufio.NewReader(os.Stdin)

	var m int
	fmt.Fscan(in, &n, &m, &C, &D)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &cost[0][i], &cost[1][i])
	}
	for i := 0; i < m; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		x--
		y--
		g[y] |= (1 << x)
	}
	fmt.Println(min(calc(0, 0, 0), calc(0, 0, 1)))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TEN(n int) int {
	if n == 0 {
		return 1
	}
	return 10 * TEN(n-1)
}

func calc(f, c, p int) int {
	if f == (1<<n)-1 {
		return 0
	}
	if used[f][c][p] {
		return dp[f][c][p]
	}
	used[f][c][p] = true
	ans := TEN(9)
	nf := f
	off := 0
	for i := 0; i < n; i++ {
		if (f & (1 << i)) != 0 {
			continue
		}
		if (g[i] & ^f) != 0 {
			continue
		}
		if cost[p][i] <= cost[1-p][i] {
			nf |= (1 << i)
			off += cost[p][i]
		}
	}
	if nf != f {
		dp[f][c][p] = calc(nf, c, p) + off
		return dp[f][c][p]
	}
	for i := 0; i < n; i++ {
		if (f & (1 << i)) != 0 {
			continue
		}
		if (g[i] & ^f) != 0 {
			continue
		}
		ans = min(ans, calc(f|(1<<i), c, p)+cost[p][i])
		ans = min(ans, calc(f|(1<<i), c+1, 1-p)+cost[1-p][i]+C*c+D)
	}
	dp[f][c][p] = ans
	return dp[f][c][p]
}
