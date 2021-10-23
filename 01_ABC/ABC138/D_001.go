package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	g   = make([][]int, 200001)
	res [200001]int
)

func dfs(i, p int) {
	for _, x := range g[i] {
		if x != p {
			res[x] += res[i]
			dfs(x, i)
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, q int
	fmt.Fscan(in, &n, &q)

	for i := 0; i < n-1; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		a--
		b--
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}

	for i := 0; i < q; i++ {
		var p, x int
		fmt.Fscan(in, &p, &x)
		p--
		res[p] += x
	}

	dfs(0, 0)

	fmt.Println(strings.Trim(fmt.Sprint(res[:n]), "[]"))
}
