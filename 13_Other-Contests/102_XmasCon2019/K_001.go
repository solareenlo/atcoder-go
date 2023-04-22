package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const N = 204000
const M = 3000000

type edge struct {
	to  int
	nxt int
}

var m, ct int
var vis [M]bool
var e [N]edge
var h, hd, ss, pre [N]int
var mp [N]map[int]int

func main() {
	in := bufio.NewReader(os.Stdin)

	n := int(2.8e6)
	for i := 2; i <= n; i++ {
		if !vis[i] {
			m++
			pre[m] = i
		}
		for j := 1; j <= m; j++ {
			u := i * pre[j]
			if u > n {
				break
			}
			vis[u] = true
			if (i % pre[j]) == 0 {
				break
			}
		}
	}

	fmt.Fscan(in, &n)
	var f [N]int
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &f[i])
		add(f[i], i)
	}

	for i := range mp {
		mp[i] = make(map[int]int)
	}
	k := make(map[int]int)
	for u := 1; u <= n; u++ {
		if f[u] != 0 {
			continue
		}
		for i := hd[u]; i > 0; i = e[i].nxt {
			v := e[i].to
			dfs(v)
			mp[u][h[v]]++
		}
		keys := make([]int, 0, len(mp[u]))
		for k := range mp[u] {
			keys = append(keys, k)
		}
		sort.Ints(keys)
		for _, i := range keys {
			k[i] ^= mp[u][i]
		}
	}
	keys := make([]int, 0, len(k))
	for i := range k {
		keys = append(keys, i)
	}
	sort.Ints(keys)
	for _, i := range keys {
		if k[i] != 0 {
			fmt.Println("Black")
			return
		}
	}
	fmt.Println("White")
}

func add(u, v int) {
	if u != 0 {
		ct++
		e[ct] = edge{v, hd[u]}
		hd[u] = ct
	}
}

func dfs(u int) {
	h[u] = 1
	ss[u] = 1
	for i := hd[u]; i > 0; i = e[i].nxt {
		v := e[i].to
		dfs(v)
		ss[u] += ss[v]
		h[u] += h[v] * pre[ss[v]^283]
	}
}
