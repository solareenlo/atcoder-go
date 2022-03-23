package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 100005

var (
	sze int
	to  = make([]int, N+N)
	nx  = make([]int, N+N)
	hd  = make([]int, N)
	p   = make([]bool, N)
)

func add(u, v int) {
	sze++
	to[sze] = v
	nx[sze] = hd[u]
	hd[u] = sze
}

func dfs(u int) {
	p[u] = true
	for i := hd[u]; i > 0; i = nx[i] {
		v := to[i]
		if !p[v] {
			dfs(v)
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	o := make(map[int]int)
	a := make([]int, N)
	x := 0
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		o[a[i]]++
		x ^= a[i]
	}
	o[x]++

	b := make([]int, N)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &b[i])
		o[b[i]]--
	}

	m := 0
	for key, val := range o {
		if val < 0 {
			fmt.Println(-1)
			return
		} else {
			m++
			o[key] = m
		}
	}

	ans := 0
	for i := 1; i <= n; i++ {
		if a[i] != b[i] {
			add(o[a[i]], o[b[i]])
			add(o[b[i]], o[a[i]])
			ans++
		}
	}

	for i := 1; i <= m; i++ {
		if !p[i] && hd[i] != 0 {
			ans++
			dfs(i)
		}
	}

	if hd[o[x]] != 0 {
		ans--
	}
	fmt.Println(ans)
}
