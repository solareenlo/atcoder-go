package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m, K int
	fmt.Fscan(in, &n, &m, &K)

	vis := [20][20]bool{}
	pos := [20][20]int{}
	for i := 1; i <= K; i++ {
		var a, b, c int
		fmt.Fscan(in, &a, &b, &c)
		a--
		b--
		vis[a][b] = true
		pos[a][b] = c
	}

	const P = 1_000_000_007
	n--
	U := 1 << n
	f := make([]int, 1<<20)
	f[0] = 1
	for p := 0; p < m; p++ {
		for i := 0; i < n; i++ {
			if !vis[p][i] || pos[p][i] != 0 {
				for j, l := 0, 1<<i; j < U; j += l << 1 {
					for k := 0; k < l; k++ {
						if f[j+k] != 0 {
							f[j+k+l-(j&-j)] += f[j+k]
						}
					}
				}
			}
			if vis[p][i] {
				l := 1 << i
				k := l
				if pos[p][i] != 0 {
					k = 0
				}
				for j := 0; j < U; j += l << 1 {
					for a := j + k; a < j+k+(l<<3)/8; a++ {
						f[a] = 0
					}
				}
			}
		}
		for i := 0; i < U; i++ {
			if f[i] >= P {
				f[i] %= P
			}
		}
	}

	ans := 0
	for i := 0; i < U; i++ {
		ans += f[i]
	}
	fmt.Println(ans % P)
}
