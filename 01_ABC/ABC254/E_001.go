package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 200010

var k, c, sum int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)

	g := make([][]int, N)
	for i := 1; i <= m; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}

	v := make([]int, N)
	var q int
	fmt.Fscan(in, &q)
	c = 0
	for q > 0 {
		sum = 0
		var x int
		fmt.Fscan(in, &x, &k)
		c++
		dfs(x, 0, v, g)
		fmt.Fprintln(out, sum)
		q--
	}
}

func dfs(x, distance int, v []int, g [][]int) {
	if distance > k {
		return
	}
	if v[x] != c {
		sum += x
	}
	v[x] = c
	for _, y := range g[x] {
		dfs(y, distance+1, v, g)
	}
}
