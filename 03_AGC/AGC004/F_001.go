package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const N = 100005

var (
	fa = make([]int, N)
	dp = make([]int, N)
	e  = make([][]int, N)
)

func findRt(x int) int {
	if x == fa[x] {
		return x
	}
	fa[x] = findRt(fa[x])
	return fa[x]
}

func merge(x, y int) {
	x = findRt(x)
	y = findRt(y)
	if x != y {
		fa[x] = y
	}
}

func DFS(u, f, fl int) {
	fa[u] = f
	dp[u] = fl
	for i := 0; i < len(e[u]); i++ {
		v := e[u][i]
		if v != f {
			DFS(v, u, -fl)
			dp[u] += dp[v]
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)
	for i := 1; i <= n; i++ {
		fa[i] = i
	}

	U, V := 1, 0
	for i := 1; i <= m; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		if findRt(u) != findRt(v) {
			merge(u, v)
			e[u] = append(e[u], v)
			e[v] = append(e[v], u)
		} else {
			U = u
			V = v
		}
	}

	DFS(U, 0, 1)

	st := make([]int, N)
	if m < n {
		if dp[U] != 0 {
			fmt.Println(-1)
			return
		}
	} else {
		for i := V; i > 0; i = fa[i] {
			st[0]++
			st[st[0]] = dp[i]
		}
		if st[0]&1 != 0 {
			if dp[U]&1 != 0 {
				fmt.Println(-1)
				return
			}
			for i := V; i > 0; i = fa[i] {
				dp[i] -= dp[U] / 2
			}
		} else {
			if dp[U] != 0 {
				fmt.Println(-1)
				return
			}
			tmp := st[1 : st[0]+1]
			sort.Ints(tmp)
			for i := V; i > 0; i = fa[i] {
				dp[i] -= st[st[0]/2]
			}
		}
	}

	ans := 0
	for i := 1; i <= n; i++ {
		ans += abs(dp[i])
	}
	fmt.Println(ans)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
