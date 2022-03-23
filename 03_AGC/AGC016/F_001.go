package main

import (
	"bufio"
	"fmt"
	"os"
)

const mod = 1_000_000_007

var (
	f = make([]int, 1<<15)
	G = [15][15]int{}
	g = [15][2]int{}
)

func DFS(s, i int) int {
	for i >= 0 && (s>>i&1) == 0 {
		i--
	}
	if i < 0 {
		return f[s]
	}
	for j := 0; j < i; j++ {
		if G[j][i] != 0 {
			g[j][0]++
		}
	}
	res := DFS(s, i-1)
	res = (res << g[i][1]) - res
	for j := 0; j < i; j++ {
		if G[j][i] != 0 {
			g[j][0]--
			g[j][1]++
		}
	}
	res += DFS(s^1<<i, i-1) << g[i][0]
	for j := 0; j < i; j++ {
		if G[j][i] != 0 {
			g[j][1]--
		}
	}
	return res % mod
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, M int
	fmt.Fscan(in, &N, &M)
	for i := 0; i < M; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		a--
		b--
		G[a][b] = 1
	}

	f[0] = 1
	for i := 0; i < N; i++ {
		for s := 0; s < 1<<i; s++ {
			if ((s | 1<<i) & 1) == ((s | 1<<i) >> 1 & 1) {
				for j := 0; j < i; j++ {
					g[j][0] = 0
					g[j][1] = G[j][i]
				}
				f[s|1<<i] = DFS(s, i)
			}
		}
	}

	s := 1
	for i := 0; i < M; i++ {
		s = s * 2 % mod
	}
	s = (s + mod - f[(1<<N)-1]) % mod
	fmt.Println(s)
}
