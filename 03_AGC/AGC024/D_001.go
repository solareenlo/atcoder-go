package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	rec int
	mx  = make([]int, 222)
	G   = make([][]int, 222)
)

func dfs(x, ff, dep int) {
	rec = max(rec, (dep+1)>>1)
	for _, y := range G[x] {
		if y != ff {
			dfs(y, x, dep+1)
		}
	}
	tmp := 0
	if ff != 0 {
		tmp = 1
	}
	mx[dep] = max(mx[dep], len(G[x])-tmp)
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	for i := 1; i < n; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		G[x] = append(G[x], n+i)
		G[n+i] = append(G[n+i], x)
		G[y] = append(G[y], n+i)
		G[n+i] = append(G[n+i], y)
	}

	res := 0
	anc := 1 << 60
	ans := 0
	for i := 1; i < n*2; i++ {
		rec = 0
		res = 1
		for j := 1; j <= n*2; j++ {
			mx[j] = 1
		}
		dfs(i, 0, 1)
		for j := 1; j <= n*2; j++ {
			res *= mx[j]
		}
		if rec < anc || rec == anc && res < ans {
			anc = rec
			ans = res
		}
	}

	fmt.Println(anc, ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
