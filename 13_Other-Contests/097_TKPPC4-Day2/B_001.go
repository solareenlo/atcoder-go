package main

import (
	"bufio"
	"fmt"
	"os"
)

var T [][]int
var a, b []bool

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)

	T = make([][]int, n)
	for i := 0; i < n-1; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		u--
		v--
		T[u] = append(T[u], v)
		T[v] = append(T[v], u)
	}

	a = make([]bool, n)
	for i := 0; i < m; i++ {
		var x int
		fmt.Fscan(in, &x)
		x--
		a[x] = true
	}

	root := -1
	for i := 0; i < n; i++ {
		if a[i] {
			root = i
		}
	}
	b = make([]bool, n)
	dfs(root, -1)

	deg := make([]int, n)
	for u := 0; u < n; u++ {
		for _, v := range T[u] {
			if b[u] && b[v] {
				deg[v]++
			}
		}
	}

	hist := make([]int, n)
	for u := 0; u < n; u++ {
		if b[u] {
			hist[deg[u]]++
		}
	}

	cnt := 0
	for _, i := range b {
		if i {
			cnt++
		}
	}
	if hist[1] == 2 && hist[2] == cnt-2 {
		fmt.Println("Yes")
	} else {
		fmt.Println("trumpet")
	}
}

func dfs(u, p int) bool {
	b[u] = a[u]
	for _, v := range T[u] {
		if v != p {
			if dfs(v, u) {
				b[u] = true
			}
		}
	}
	return b[u]
}
