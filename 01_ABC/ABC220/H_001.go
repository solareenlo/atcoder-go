package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	n   int
	cnt = [1 << 20][2]int{}
	to  = [45]int{}
	to2 = [45]int{}
)

func popcount(x int) int {
	res := 0
	for ; x > 0; x ^= x & -x {
		res ^= 1
	}
	return res
}

func dfs(x, now, s, t int) {
	if x == n {
		cnt[s][t]++
		return
	}
	dfs(x+1, now, s, t)
	dfs(x+1, 1<<x|now, s^to[x], t^(popcount(to2[x]&now)))
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var m int
	fmt.Fscan(in, &n, &m)
	lmt := n >> 1
	for i := 1; i <= m; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		x--
		y--
		if x < lmt {
			if y < x || y >= lmt {
				to[y] |= 1 << x
			}
		} else {
			to2[y] |= 1 << x
		}
		if y < lmt {
			if x < y || x >= lmt {
				to[x] |= 1 << y
			}
		} else {
			to2[x] |= 1 << x
		}
	}

	dfs(lmt, 0, 0, 0)

	g := [1 << 20][2]int{}
	for i := 0; i < lmt; i++ {
		for j := 0; j < (1 << lmt); j++ {
			for k := 0; k < 2; k++ {
				g[j][k] = cnt[j][k]
				cnt[j][k] = 0
			}
		}
		for s := 0; s < (1 << lmt); s++ {
			d := s >> i & 1
			for t := 0; t < 2; t++ {
				cnt[s^to[i]][t] += g[s][t]
				cnt[d<<i^s][t] += g[s][t]
			}
		}
	}
	ans := 0
	m &= 1
	for s := 0; s < (1 << lmt); s++ {
		for t := 0; t < 2; t++ {
			if (popcount(s) ^ t) == m {
				ans += cnt[s][t]
			}
		}
	}
	fmt.Println(ans)
}
