package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var x, c, v [18]int
	var xs [19]int
	var cs, vs, dp [1 << 18]int
	var f [1 << 18][19]int

	var N int
	fmt.Fscan(in, &N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &x[i])
		xs[i+1] = xs[i] + x[i]
	}
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &c[i])
	}
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &v[i])
	}
	for i := 0; i < 1<<N; i++ {
		for j := 0; j < N; j++ {
			if ((i >> j) & 1) != 0 {
				cs[i] += c[j]
				vs[i] += v[j]
			}
		}
	}
	for i := 1<<N - 1; i >= 0; i-- {
		for j := 0; j <= N; j++ {
			if cs[^i&((1<<N)-1)] <= xs[j] {
				f[i][j] = vs[^i&((1<<N)-1)]
			} else {
				for k := 0; k < N; k++ {
					f[i][j] = max(f[i][j], f[i|(1<<k)][j])
				}
			}
		}
	}
	for i := (1 << N) - 1 - 1; i >= 0; i-- {
		t := int(9e18)
		for j := 0; j < N; j++ {
			if (i>>j)&1 == 0 {
				t = min(t, dp[i|(1<<j)])
			}
		}
		dp[i] = max(f[i][popcount(uint32(i))], t)
	}
	fmt.Println(dp[0])
}

func popcount(n uint32) int {
	return bits.OnesCount32(n)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
