package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const N = 200005

type E struct {
	u, v, w int
}

var s, cnt int
var f, g [N + 5]int

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	var d [N + 5]int
	var e [N + 5]E
	for i := 1; i < 2*n-1; i++ {
		var u, v, w int
		fmt.Fscan(in, &u, &v, &w)
		if i >= n {
			if u > 1 {
				u += n
			}
			if v > 1 {
				v += n
			}
		}
		d[u]++
		d[v]++
		e[i] = E{u, v, w}
	}

	for i := 1; i <= 2*n; i++ {
		if d[i] == 1 {
			s++
			g[i] = 1
		}
		f[i] = i
	}
	cnt = s
	s /= 2
	tmp := e[1 : 2*n-1]
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i].w > tmp[j].w
	})
	ans := 0
	for i := 1; i < 2*n-1; i++ {
		if !merge(e[i].u, e[i].v) {
			ans += e[i].w
		}
	}
	fmt.Println(ans)
}

func merge(u, v int) bool {
	u = find(u)
	v = find(v)
	if u == v {
		return false
	}
	t := g[u] + g[v] - (g[u] | g[v])
	if cnt-t <= s {
		return false
	}
	cnt -= t
	f[u] = v
	g[v] |= g[u]
	return true
}

func find(x int) int {
	if f[x] == x {
		return f[x]
	}
	f[x] = find(f[x])
	return f[x]
}
