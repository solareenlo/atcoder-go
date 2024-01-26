package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const MOD = 998244353
	const N = 200200

	var banc, bans [N]int

	var n, m int
	fmt.Fscan(in, &n, &m)
	var g [N][]int
	for i := 1; i <= m; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}
	dist := make([]int, N)
	for i := range dist {
		dist[i] = 1061109567
	}
	dist[1] = 0
	pc := make([]int, N)
	pc[1] = 1
	for d := 0; ; d++ {
		cnt, sum := 0, 0
		for i := 1; i <= n; i++ {
			banc[i] = 0
			bans[i] = 0
		}
		for i := 1; i <= n; i++ {
			if dist[i] == d {
				cnt++
				sum = (sum + pc[i]) % MOD
				for _, v := range g[i] {
					banc[v]++
					bans[v] = (bans[v] + pc[i]) % MOD
				}
			}
		}
		if cnt == 0 {
			break
		}
		for i := 1; i <= n; i++ {
			if dist[i] > d {
				if banc[i] < cnt {
					dist[i] = d + 1
					pc[i] = (sum - bans[i] + MOD) % MOD
				}
			}
		}
	}
	if dist[n] >= 1e9 {
		fmt.Println(-1)
	} else {
		fmt.Println(pc[n])
	}
}
