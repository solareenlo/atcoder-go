package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 200010

var vis [N]bool
var e [N][]tuple
var x, y [N]int

func dfs(a int) {
	vis[a] = true
	for _, tmp := range e[a] {
		b, dx, dy := tmp.x, tmp.y, tmp.z
		if vis[b] {
			continue
		}
		x[b] = x[a] + dx
		y[b] = y[a] + dy
		dfs(b)
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)
	for ; m > 0; m-- {
		var a, b, x, y int
		fmt.Fscan(in, &a, &b, &x, &y)
		e[a] = append(e[a], tuple{b, x, y})
		e[b] = append(e[b], tuple{a, -x, -y})
	}
	dfs(1)
	for i := 1; i <= n; i++ {
		if !vis[i] {
			fmt.Fprintln(out, "undecidable")
		} else {
			fmt.Fprintln(out, x[i], y[i])
		}
	}
}

type tuple struct {
	x, y, z int
}
