package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 100_004

type pair struct{ x, y int }

var (
	e   = make([][]pair, N)
	k   int
	cnt int
	f   = map[int]int{}
)

func dfs(x, p, v int) {
	cnt += f[k^v]
	f[v]++
	for _, i := range e[x] {
		if i.x != p {
			dfs(i.x, x, v^i.y)
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n, &k)

	for i := 1; i < n; i++ {
		var x, y, c int
		fmt.Fscan(in, &x, &y, &c)
		e[x] = append(e[x], pair{y, c})
		e[y] = append(e[y], pair{x, c})
	}

	dfs(1, 0, 0)

	fmt.Println(cnt)
}
