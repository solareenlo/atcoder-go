package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 1e5 + 10

var (
	n   int
	vis = make([]bool, N)
	c   = make([]int, 0)
	a   = make([]int, N)
	deg = make([]int, N)
)

func dfs(nw, lst int) {
	if nw >= n {
		for _, v := range c {
			fmt.Print(v, " ")
		}
		fmt.Println()
		os.Exit(0)
	}
	if len(c) != 0 {
		u := a[c[len(c)-1]]
		if !vis[u] && nw+deg[u]+1 == n {
			return
		}
	}
	for lst <= n && vis[lst] {
		lst++
	}
	for i := lst; i <= n; i++ {
		if !vis[i] {
			if nw != 0 && i == a[c[nw-1]] {
				continue
			}
			vis[i] = true
			c = append(c, i)
			deg[a[i]]--
			dfs(nw+1, lst)
			vis[i] = false
			c = c[:len(c)-1]
			deg[a[i]]++
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		deg[a[i]]++
	}
	dfs(0, 1)
	fmt.Println(-1)
}
