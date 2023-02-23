package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 5145
	const mod = 1000000007

	var n, m int
	fmt.Fscan(in, &n, &m)
	var f [2][N]int
	f[0][3] = 1
	M := 1 << n
	var g [N][N]int
	for j := 1; j < M; j++ {
		for k := 1; k < M*M; k++ {
			p, nw, pd, c := 0, 0, 0, 0
			ls := 2
			for i := 0; i <= n; i++ {
				if (j >> i & 1) != 0 {
					u := (k >> (2 * i)) & 3
					if u == 2 && c == 3 || u == 3 && c == 2 {
						for o := 0; o < i; o++ {
							if ((nw >> (2 * o)) & 3) == ls {
								nw += (3 - ls) << (2 * o)
							}
						}
						ls = 3
					} else if u == 2 {
						u = ls
					}
					c = max(c, u)
				} else {
					if p != i && pd == 0 && c < 2 {
						c = 2
						if ls != 3 {
							ls = 1
						}
					}
					for o := p; o < i; o++ {
						nw ^= c << (o * 2)
					}
					if c == 2 {
						pd = 1
					}
					c = 0
					p = i + 1
				}
			}
			g[j][k] = nw
		}
	}
	x := 0
	for l := 1; l <= m; l++ {
		x ^= 1
		for i := range f[x] {
			f[x][i] = 0
		}
		for j := 1; j < M; j++ {
			for k := 1; k < M*M; k++ {
				if f[x^1][k] != 0 {
					f[x][g[j][k]] += f[x^1][k]
					f[x][g[j][k]] %= mod
				}
			}
		}
	}
	res := 0
	for i := 3 << (2*n - 2); i < 1<<(2*n); i++ {
		res = (res + f[x][i]) % mod
	}
	fmt.Println(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
