package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const N = 300003

var d, hd, f, a, g [N]int
var nxt, ver, val [N << 1]int
var tot int

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &d[i])
	}

	for i := 1; i < n; i++ {
		var u, v, w int
		fmt.Fscan(in, &u, &v, &w)
		add(u, v, w)
		add(v, u, w)
	}

	dfs(1, 0)
	fmt.Println(f[1])
}

func add(u, v, w int) {
	tot++
	nxt[tot] = hd[u]
	hd[u] = tot
	ver[tot] = v
	val[tot] = w
}

func dfs(u, fa int) {
	sum, rk := 0, 0
	for i := hd[u]; i > 0; i = nxt[i] {
		v := ver[i]
		if v == fa {
			continue
		}
		dfs(v, u)
		sum += f[v]
	}
	for i := hd[u]; i > 0; i = nxt[i] {
		v := ver[i]
		if v == fa {
			continue
		}
		if d[v] != 0 {
			a[rk] = f[v] - g[v] - val[i]
			rk++
		}
	}
	tmp := a[:rk]
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i] < tmp[j]
	})
	f[u] = sum
	g[u] = sum
	for i := 0; a[i] < 0 && i < d[u] && i < rk; i++ {
		f[u] -= a[i]
		if i < d[u]-1 {
			g[u] -= a[i]
		}
	}
}
