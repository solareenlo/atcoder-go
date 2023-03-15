package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)

	adj := make([][]int, n)
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		a--
		b--
		adj[b] = append(adj[b], a)
	}

	reach := make([][]int, n)
	vis := make([]int, n)
	for i := range vis {
		vis[i] = -1
	}
	for i := 0; i < n; i++ {
		var dfs func(int)
		dfs = func(x int) {
			if vis[x] == i {
				return
			}
			reach[i] = append(reach[i], x)
			vis[x] = i
			if len(reach[i]) > 3 {
				return
			}
			for _, y := range adj[x] {
				dfs(y)
				if len(reach[i]) > 3 {
					return
				}
			}
		}
		dfs(i)
	}

	var cnt [5]int
	for i := 0; i < n; i++ {
		cnt[len(reach[i])]++
	}

	g := make([]int, n)
	for i := 0; i < n; i++ {
		if len(reach[i]) == 2 {
			g[reach[i][1]]++
		}
	}

	ans := 0
	for i := 0; i < n; i++ {
		if len(reach[i]) == 3 {
			x := reach[i][1]
			y := reach[i][2]
			if len(reach[x]) == 1 && len(reach[y]) == 1 {
				ans += 2
			} else {
				ans++
			}
		}
		if len(reach[i]) == 2 {
			ans += cnt[1] - 1
			ans += cnt[1] + g[reach[i][1]] - 2
		}
		if len(reach[i]) == 1 {
			ans += (cnt[1] - 1) * (cnt[1] - 2)
			ans += cnt[2]
			ans -= g[i]
		}
	}
	fmt.Println(ans)
}
