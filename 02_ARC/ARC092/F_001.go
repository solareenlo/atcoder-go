package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 1007
const M = 200007

var (
	c = [2][N][N]int{}
	G = make([][]int, N)
	a = [M]int{}
	b = [M]int{}
)

func dfs(x, st, ptr, typ int) {
	if c[typ][st][x] == 0 {
		c[typ][st][x] = ptr
	} else {
		return
	}
	for _, y := range G[x] {
		if c[typ][st][y] == 0 {
			dfs(y, st, ptr, typ)
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)

	for i := 1; i <= m; i++ {
		fmt.Fscan(in, &a[i], &b[i])
		G[a[i]] = append(G[a[i]], b[i])
	}

	for x := 1; x <= n; x++ {
		c[0][x][x] = 1
		c[1][x][x] = 1
		for _, y := range G[x] {
			dfs(y, x, y, 0)
		}
		G[x] = reverseOrderInt(G[x])
		for _, y := range G[x] {
			dfs(y, x, y, 1)
		}
	}

	for i := 1; i <= m; i++ {
		flag := 0
		if c[0][b[i]][a[i]] > 0 {
			flag = 1
		}
		if c[0][a[i]][b[i]] != c[1][a[i]][b[i]] {
			flag ^= 1
		}
		if flag != 0 {
			fmt.Fprintln(out, "diff")
		} else {
			fmt.Fprintln(out, "same")
		}
	}
}

func reverseOrderInt(a []int) []int {
	n := len(a)
	res := make([]int, n)
	n = copy(res, a)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}
