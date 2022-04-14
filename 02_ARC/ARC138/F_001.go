package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const MAX = 30
const Mod = 1_000_000_007

var (
	n   int
	dp  = [MAX + 5][MAX + 5][MAX + 5][MAX + 5]int{}
	p   = [MAX + 5]int{}
	pos = [MAX + 5]int{}
)

func Solv(xl, xr, yl, yr int) int {
	if dp[xl][xr][yl][yr] != -1 {
		return dp[xl][xr][yl][yr]
	}
	sz := 0
	a := make([]int, MAX+5)
	b := make([]int, MAX+5)
	for i := xl; i <= xr; i++ {
		if p[i] >= yl && p[i] <= yr {
			sz++
			a[sz] = i
			b[sz] = p[i]
		}
	}
	if sz <= 1 {
		dp[xl][xr][yl][yr] = 1
		return 1
	}
	tmp := a[1 : sz+1]
	sort.Ints(tmp)
	tmp = b[1 : sz+1]
	sort.Ints(tmp)
	x := [MAX + 5]int{}
	y := [MAX + 5]int{}
	dp[xl][xr][yl][yr] = 0
	for i := 1; i < sz; i++ {
		x[i] = Solv(xl, a[i], yl, yr)
		for j := 1; j < i; j++ {
			x[i] = (x[i] - x[j]*Solv(a[j+1], a[i], yl, yr)%Mod + Mod) % Mod
		}
		y[i] = Solv(xl, xr, yl, b[i])
		for j := 1; j < i; j++ {
			y[i] = (y[i] - y[j]*Solv(xl, xr, b[j+1], b[i])%Mod + Mod) % Mod
		}
		Mn := n
		for j := i + 1; j <= sz; j++ {
			Mn = min(Mn, p[a[j]])
		}
		for j := 1; b[j] < Mn; j++ {
			x[i] = (x[i] - y[j]*Solv(xl, a[i], b[j+1], yr)%Mod + Mod) % Mod
		}
		Mn = n
		for j := i + 1; j <= sz; j++ {
			Mn = min(Mn, pos[b[j]])
		}
		for j := 1; a[j] < Mn; j++ {
			y[i] = (y[i] - x[j]*Solv(a[j+1], xr, yl, b[i])%Mod + Mod) % Mod
		}
		dp[xl][xr][yl][yr] = (dp[xl][xr][yl][yr] + x[i]*Solv(a[i+1], xr, yl, yr)%Mod) % Mod
		dp[xl][xr][yl][yr] = (dp[xl][xr][yl][yr] + y[i]*Solv(xl, xr, b[i+1], yr)%Mod) % Mod
	}
	return dp[xl][xr][yl][yr]
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &p[i])
		pos[p[i]] = i
	}

	for i := range dp {
		for j := range dp[i] {
			for k := range dp[i][j] {
				for l := range dp[i][j][k] {
					dp[i][j][k][l] = -1
				}
			}
		}
	}
	fmt.Println(Solv(1, n, 1, n))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
