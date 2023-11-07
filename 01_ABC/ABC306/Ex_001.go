package main

import (
	"bufio"
	"fmt"
	"os"
)

var fa [11 + 45 + 14]int

func getfa(pos int) int {
	if fa[pos] != pos {
		fa[pos] = getfa(fa[pos])
	}
	return fa[pos]
}

func count(pos int) int {
	cnt := 0
	for pos != 0 {
		cnt++
		pos = pos & (pos - 1)
	}
	return cnt
}

func main() {
	in := bufio.NewReader(os.Stdin)

	const MOD = 998244353

	type pair struct {
		x, y int
	}
	var e [114 + 514]pair

	var f, comp [1 << 17]int
	f[0] = 1

	var n, m int
	fmt.Fscan(in, &n, &m)
	for i := 1; i <= m; i++ {
		fmt.Fscan(in, &e[i].x, &e[i].y)
	}
	for i := 0; i < 1<<n; i++ {
		comp[i] = count(i)
		for j := 1; j <= n; j++ {
			fa[j] = j
		}
		for j := 1; j <= m; j++ {
			if (((i >> (e[j].x - 1)) & 1) != 0) && ((i>>(e[j].y-1))&1) != 0 {
				if getfa(e[j].x) != getfa(e[j].y) {
					fa[getfa(e[j].x)] = getfa(e[j].y)
					comp[i]--
				}
			}
		}
	}
	for i := 1; i < (1 << n); i++ {
		for j := i; j != 0; j = i & (j - 1) {
			if (comp[j] & 1) != 0 {
				f[i] += f[i^j]
			} else {
				f[i] -= f[i^j]
			}
		}
		f[i] %= MOD
	}
	fmt.Println((f[(1<<n)-1] + MOD) % MOD)
}
