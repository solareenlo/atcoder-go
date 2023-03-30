package main

import (
	"bufio"
	"fmt"
	"os"
)

const N int = 150001

var (
	i, j, t, n int
	d          [N]int
	f          [N]float64
	g          [N]float64
	fa         [N]int
	en, nxt    [N + N]int
)

func dfs1(u, p int) {
	for i := fa[u]; i > 0; i = nxt[i] {
		if en[i] != p {
			d[u]++
			dfs1(en[i], u)
			f[u] += f[en[i]]
		}
	}
	if d[u] != 0 {
		f[u] = f[u]/float64(d[u]) + 1
	}
}

func dfs2(u, p int) {
	for i, v := fa[u], 0; i > 0; i = nxt[i] {
		if v = en[i]; v != p {
			if p != 0 {
				g[v] = (f[u]*float64(d[u]) - f[v] - 1 + g[u]) / float64(d[u])
			} else if d[u] > 1 {
				g[v] = (f[u]*float64(d[u]) - f[v] - 1) / float64(d[u]-1)
			}
			g[v]++
			dfs2(v, u)
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	fmt.Fscan(in, &n)
	for {
		t++
		if t >= n+n-1 {
			break
		}
		fmt.Fscan(in, &i, &j)
		en[t] = j
		nxt[t] = fa[i]
		fa[i] = t
		t++
		en[t] = i
		nxt[t] = fa[j]
		fa[j] = t
	}
	dfs1(1, 0)
	dfs2(1, 0)
	for i = 1; i <= n; i++ {
		if g[i] != 0 {
			fmt.Fprintf(out, "%.6f\n", (f[i]*float64(d[i])+g[i])/float64(d[i]+1))
		} else {
			fmt.Fprintf(out, "%.6f\n", f[i])
		}
	}
}
