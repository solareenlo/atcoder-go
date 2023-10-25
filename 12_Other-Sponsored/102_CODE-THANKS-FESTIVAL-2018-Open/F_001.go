package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var dis [303]int
var G [303][]int
var used [303]bool
var cnt []int

func Init(u, d int) {
	dis[u] = d
	for _, v := range G[u] {
		Init(v, d+1)
	}
}

func dfs(u int) {
	if !used[u] {
		cnt = append(cnt, dis[u])
		for _, v := range G[u] {
			dfs(v)
		}
	}
}

func Make(u int) {
	used[u] = true
	for _, v := range G[u] {
		Make(v)
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m, K int
	fmt.Fscan(in, &n, &m, &K)
	b := 0
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(in, &a)
		b++
		G[a] = append(G[a], b)
	}
	root := G[0][0]
	Init(root, 1)
	ans := make([]int, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; {
			j++
			if used[j] {
				continue
			}
			used[j] = true
			cnt = make([]int, 0)
			dfs(root)
			sort.Ints(cnt)
			if len(cnt) >= m-i-1 {
				var a, b, k int
				for a, b, k = 0, 0, 0; k < m-i-1; k++ {
					a += cnt[k]
					b += cnt[len(cnt)-k-1]
				}
				if a <= K-dis[j] && K-dis[j] <= b {
					ans = append(ans, j)
					Make(j)
					K -= dis[j]
					break
				}
			}
			used[j] = false
		}
	}
	if len(ans) < m {
		ans = []int{-1}
	}
	for _, v := range ans {
		fmt.Print(v, " ")
	}
}
