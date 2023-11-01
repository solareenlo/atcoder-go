package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const INF = int(1e18)
	const N = 505

	var n int
	fmt.Fscan(in, &n, &MOD)
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	d := make([]int, N)
	for i := range d {
		d[i] = -INF
	}
	d[1] = 0
	ok := make([]bool, N)
	for i := 1; i <= n; i++ {
		u := 0
		for j := 1; j <= n; j++ {
			if !ok[j] && d[j] > d[u] {
				u = j
			}
		}
		ok[u] = true
		for v := 1; v <= n; v++ {
			if !ok[v] {
				d[v] = max(d[v], (powMod(a[u], a[v])+powMod(a[v], a[u]))%MOD)
			}
		}
	}
	ans := 0
	for i := 2; i <= n; i++ {
		ans += d[i]
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

var MOD int

func powMod(a, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = res * a % MOD
		}
		a = a * a % MOD
		n /= 2
	}
	return res
}
