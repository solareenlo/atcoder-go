package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 1000005

var n, m int
var f, dt []int
var vis, vv []bool
var e [][]int

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n, &m)

	f = make([]int, N)
	for i := 1; i <= n*2; i++ {
		f[i] = i
	}
	x := make([]int, N)
	y := make([]int, N)
	e = make([][]int, N)
	for i := 1; i <= m; i++ {
		fmt.Fscan(in, &x[i], &y[i])
		e[x[i]] = append(e[x[i]], y[i])
	}
	vis = make([]bool, N)
	vv = make([]bool, N)
	dfs(1)
	if !vis[n] {
		fmt.Println(-1)
		return
	}
	dt = make([]int, N)
	for i := 1; i <= m; i++ {
		if vis[x[i]] && vis[y[i]] {
			f[find(x[i]*2)] = find(y[i]*2 - 1)
		}
	}

	a := make([]int, N)
	ans := 0
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		if !vis[i] {
			continue
		}
		if a[i] == -1 {
			continue
		}
		if find(i*2-1) != find(i*2) {
			dt[f[i*2-1]] = a[i] - dt[i*2-1]
			f[f[i*2-1]] = i * 2
		} else {
			ans = gcd(ans, abs(dt[i*2]+a[i]-dt[i*2-1]))
		}
	}
	if find(1) == (n * 2) {
		ans = gcd(ans, dt[1])
	}
	if ans == 0 {
		ans--
	}
	fmt.Println(ans)
}

func dfs(u int) {
	vv[u] = true
	if u == n {
		vis[u] = true
	}
	for _, v := range e[u] {
		if !vv[v] {
			dfs(v)
		}
		if vis[v] {
			vis[u] = true
		}
	}
}

func find(x int) int {
	if f[x] == x {
		return x
	}
	find(f[x])
	dt[x] = dt[x] + dt[f[x]]
	f[x] = f[f[x]]
	return f[x]
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
