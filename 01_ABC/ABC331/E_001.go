package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 100005

	var a [N]int

	var n, m, k int
	fmt.Fscan(in, &n, &m, &k)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	b := make([]pair, m+1)
	for i := 1; i <= m; i++ {
		fmt.Fscan(in, &(b[i].x))
		b[i].y = i
	}
	vis := make([]map[int]bool, N)
	for i := range vis {
		vis[i] = make(map[int]bool)
	}
	for i := 1; i <= k; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		vis[u][v] = true
	}
	sortPair(b[1:])
	ans := 0
	for i := 1; i <= n; i++ {
		p := m
		for vis[i][b[p].y] {
			p--
		}
		if p != 0 {
			ans = max(ans, a[i]+b[p].x)
		}
	}
	fmt.Println(ans)
}

type pair struct {
	x, y int
}

func sortPair(tmp []pair) {
	sort.Slice(tmp, func(i, j int) bool {
		if tmp[i].x == tmp[j].x {
			return tmp[i].y < tmp[j].y
		}
		return tmp[i].x < tmp[j].x
	})
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
