package main

import (
	"bufio"
	"fmt"
	"os"
)

const M = 1000005

var ans, ed, to [M]int

func dfs(x int) {
	if ans[x] != 0 || ed[x] != 0 {
		return
	}
	dfs(to[x])
	ans[x] = ans[to[x]]
	ed[x] = -ed[to[x]]
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var a [M]int

	var n, m int
	fmt.Fscan(in, &n, &m)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	l, r, p := 1, M-5, 0
	for i := 1; i <= m; i++ {
		var x int
		fmt.Fscan(in, &x)
		if p < l {
			p += x
		} else {
			p -= x
		}
		if p < l || p > r {
			continue
		}
		ans[p] = i
		L, R := p-l, r-p
		if L >= R {
			for j := p + 1; j <= r; j++ {
				to[j] = 2*p - j
			}
			r = p - 1
		} else {
			for j := l; j < p; j++ {
				to[j] = 2*p - j
			}
			l = p + 1
		}
	}
	for i := l; i <= r; i++ {
		ed[i] = i - p
	}
	for i := 1; i <= n; i++ {
		x := a[i]
		dfs(x)
		if ans[x] != 0 {
			fmt.Fprintf(out, "Yes %d\n", ans[x])
		} else {
			fmt.Fprintf(out, "No %d\n", ed[x])
		}
	}
}
