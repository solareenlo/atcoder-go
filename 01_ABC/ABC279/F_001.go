package main

import (
	"bufio"
	"fmt"
	"os"
)

const maxn = 300005

var fa [maxn << 1]int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, q int
	fmt.Fscan(in, &n, &q)

	var N [maxn]int
	var ans, p [maxn << 1]int
	for i := 1; i <= n; i++ {
		ans[i] = i
		fa[i] = i
		p[i] = i
		N[i] = i
	}

	tot, t := n, n
	for q > 0 {
		q--
		var op int
		fmt.Fscan(in, &op)
		switch op {
		case 1:
			var x, y int
			fmt.Fscan(in, &x, &y)
			fa[N[y]] = fa[N[x]]
			tot++
			N[y] = tot
			ans[tot] = y
			fa[tot] = tot
		case 2:
			var x int
			fmt.Fscan(in, &x)
			t++
			p[t] = N[x]
		default:
			var x int
			fmt.Fscan(in, &x)
			fmt.Fprintln(out, ans[Find(p[x])])
		}
	}
}

func Find(x int) int {
	if fa[x] == x {
		fa[x] = x
	} else {
		fa[x] = Find(fa[x])
	}
	return fa[x]
}
