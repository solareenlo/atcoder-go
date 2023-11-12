package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	const N = 1 << 15

	var dp, S [N]int
	var rev [16]int
	var vis [N]bool
	var mps [15][15]int

	var n, m int
	fmt.Fscan(in, &n, &m)
	for i, j := n, 0; i >= 0; {
		rev[n-i] = j
		dp[j] = 1
		vis[j] = true
		i--
		if i >= 0 {
			j |= (1 << i)
		}
	}
	for i := 0; i < 1<<n; i++ {
		S[i] = i
		for a := 0; a < n; a++ {
			for b := a + 1; b < n; b++ {
				if ((i >> a) & 1) != 0 {
					if ((^i >> b) & 1) != 0 {
						mps[a][b]++
					}
				}
			}
		}
	}
	for m > 0 {
		m--
		var x, y int
		fmt.Fscan(in, &x, &y)
		x = (dp[0] + x) % n
		y = (dp[0]*2 + y) % n
		if x > y {
			x, y = y, x
		}
		if mps[x][y] != 0 {
			flag := false
			for i := 0; i < 1<<n; i++ {
				if ((S[i]>>x)&1) != 0 && ((^S[i]>>y)&1) != 0 {
					for a := 0; a < y; a++ {
						if ((S[i] >> a) & 1) != 0 {
							mps[a][y]--
						}
					}
					for b := n - 1; b > x; b-- {
						if ((^S[i] >> b) & 1) != 0 {
							mps[x][b]--
						}
					}
					mps[x][y]++
					S[i] ^= (1 << x) ^ (1 << y)
					for a := 0; a < x; a++ {
						if ((S[i] >> a) & 1) != 0 {
							mps[a][x]++
						}
					}
					for b := n - 1; b > y; b-- {
						if ((^S[i] >> b) & 1) != 0 {
							mps[y][b]++
						}
					}
					if rev[popcount(uint32(i))] == S[i] {
						vis[i] = true
						flag = true
					}
				}
			}
			if flag {
				for i := (1 << n) - 2; i >= 0; i-- {
					if vis[i] {
						dp[i] = 0
						for j := 0; j < n; j++ {
							if (^i>>j)&1 != 0 {
								dp[i] += dp[i|1<<j]
							}
						}
					}
				}
			}
		}
		fmt.Fprintln(out, dp[0])
	}
}

func popcount(n uint32) int {
	return bits.OnesCount32(n)
}
