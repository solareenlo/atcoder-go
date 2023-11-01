package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func c2(n int) int { return n * (n - 1) / 2 }

func mask(i int) int {
	return (1 << i) - 1
}

func main() {
	in := bufio.NewReader(os.Stdin)

	const MOD = 998244353
	const nmax = 13

	var sub, add [1 << nmax][nmax]int
	var bsum [1 << nmax]int
	var dp [2][1 << nmax][nmax * (nmax - 1) / 2]int
	var work [1 << nmax][nmax * (nmax - 1) / 2][2]int

	var n int
	fmt.Fscan(in, &n)
	use := make([]int, c2(n)+1)
	for i := 0; i < n-1; i++ {
		var tmp int
		fmt.Fscan(in, &tmp)
		use[tmp]++
	}

	for i := 0; i < 1<<nmax; i++ {
		for j := 0; j < nmax; j++ {
			sub[i][j] = -1
			add[i][j] = -1
		}
	}
	for bit := 1; bit < 1<<n; bit++ {
		p, cur := -1, 0
		for i := 0; i < n; i++ {
			if (bit & (1 << i)) == 0 {
				p = i
			} else {
				if p >= 0 {
					sub[bit][cur] = bit ^ (3 << p)
				}
				cur++
			}
		}
		if bit < (1 << (n - 1)) {
			cnt := 0
			for i := 0; i < n; i++ {
				if (bit & (1 << i)) == 0 {
					add[bit][cnt] = (bit & mask(i)) | 1<<i | (bit>>i)<<(i+1)
					cnt++
				}
			}
		}
		cnt := 0
		for i := 0; i < n; i++ {
			if (bit & (1 << i)) == 0 {
				cnt++
			} else {
				bsum[bit] += cnt
			}
		}
	}

	tot := c2(n)

	cur, used := 0, 1
	dp[cur][1<<(n-1)][c2(n-1)] = 1
	for step := 0; step < tot+1; step++ {
		for i := 0; i < use[step]; i++ {
			for bit := 1; bit < 1<<n; bit++ {
				Len := popcount(bit)
				b := bsum[bit]
				cmax := min(tot-step-b, c2(n-Len))
				for c := 0; c < cmax+1; c++ {
					dp[cur][bit][c] = dp[cur][bit][c] * (Len - used) % MOD
				}
			}
			used++
		}
		if step == tot {
			break
		}
		nx := cur ^ 1
		for i := 0; i < 1<<nmax; i++ {
			for j := 0; j < nmax*(nmax-1)/2; j++ {
				dp[nx][i][j] = 0
			}
		}
		for bit := 1; bit < 1<<n; bit++ {
			Len := popcount(bit)
			b := bsum[bit]
			cmax := min(tot-step-b, c2(n-Len))
			for c := 0; c < cmax+1; c++ {
				a := tot - step - b - c
				// use a
				dp[nx][bit][c] = (dp[nx][bit][c] + dp[cur][bit][c]*a%MOD) % MOD
				// use c
				if c > 0 {
					dp[nx][bit][c-1] = (dp[nx][bit][c-1] + dp[cur][bit][c]*(c2(n-Len)-c+1)%MOD) % MOD
				}
				// use b
				work[bit][c][0] = dp[cur][bit][c]
				work[bit][c][1] = 0
			}
		}
		// use b
		for pos := 0; pos < n; pos++ {
			for bit := 1; bit < 1<<n; bit++ {
				Len := popcount(bit)
				b := bsum[bit]
				cmax := min(tot-step-b, c2(n-Len))
				if pos < Len {
					to := sub[bit][pos]
					if to >= 0 {
						for c := 0; c < cmax+1; c++ {
							work[to][c][0] = (work[to][c][0] + work[bit][c][0]) % MOD
							work[to][c][1] = (work[to][c][1] + work[bit][c][0]) % MOD
							work[to][c][1] = (work[to][c][1] + work[bit][c][1]) % MOD
						}
					}
				}
			}
		}
		for bit := 1; bit < 1<<(n-1); bit++ {
			Len := popcount(bit)
			b := bsum[bit]
			cmax := min(tot-step-b, c2(n-Len))
			for c := 0; c < cmax+1; c++ {
				dmax := min(c, n-Len-1)
				for d := 0; d < dmax+1; d++ {
					to := add[bit][d]
					dp[nx][to][c-d] = (dp[nx][to][c-d] + work[bit][c][1]) % MOD
				}
			}
		}
		cur, nx = nx, cur
	}
	fmt.Println(dp[cur][mask(n)][0])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func popcount(n int) int {
	return bits.OnesCount64(uint64(n))
}
