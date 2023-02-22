package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const MOD = 1000000007

	type pair struct {
		x, y int
	}

	var X, Y, N int
	fmt.Fscan(in, &X, &Y, &N)

	var small [12][12]int
	mp := make(map[pair]int)

	choose := func(N, K int) int {
		ans := 1
		for i := 0; i < K; i++ {
			ans = ans * (N - i) % MOD
		}
		for i := 1; i <= K; i++ {
			var inv int
			for inv = 1; ; inv += MOD {
				if inv%i == 0 {
					break
				}
			}
			inv /= i
			ans = ans * inv % MOD
		}
		return ans
	}

	var FUNC func(int, int, int) int
	FUNC = func(N, X, Y int) int {
		if X > Y {
			X, Y = Y, X
		}
		if X <= 2*N && Y <= 2*N {
			return small[X][Y]
		}
		p := pair{X, Y}
		if _, ok := mp[p]; ok {
			return mp[p]
		}
		var d [10][10]int
		for i := 0; i < N+1; i++ {
			d[0][i] = FUNC(N, X, N+i)
		}
		for i := 1; i <= N; i++ {
			for j := 0; j < N-i+1; j++ {
				d[i][j] = (d[i-1][j+1] - d[i-1][j] + MOD) % MOD
			}
		}
		ans := 0
		for i := 0; i < N+1; i++ {
			ans = (ans + d[i][0]*choose(Y-N, i)) % MOD
		}
		mp[p] = ans
		return mp[p]
	}

	var dp [12][12][(1 << 12)][6]int
	brute := func(N, Y int) {
		X := 2 * N
		for i := 0; i < X+1; i++ {
			for j := 0; j < Y; j++ {
				for mask := 0; mask < (1 << Y); mask++ {
					for k := 0; k < N+1; k++ {
						dp[i][j][mask][k] = 0
					}
				}
			}
		}
		dp[0][0][0][0] = 1
		for i := 0; i < X; i++ {
			for j := 0; j < Y; j++ {
				i2 := i
				j2 := j + 1
				if j2 == Y {
					i2++
					j2 = 0
				}
				for mask := 0; mask < (1 << Y); mask++ {
					for k := 0; k < N+1; k++ {
						if dp[i][j][mask][k] != 0 {
							mask2 := mask
							if (mask2 & (1 << j)) != 0 {
								mask2 ^= (1 << j)
							}
							dp[i2][j2][mask2][k] = (dp[i2][j2][mask2][k] + dp[i][j][mask][k]) % MOD
							if k < N && (mask&(1<<j)) == 0 {
								if i+1 < X {
									dp[i2][j2][mask2|(1<<j)][k+1] = (dp[i2][j2][mask2|(1<<j)][k+1] + dp[i][j][mask][k]) % MOD
								}
								if j+1 < Y && (mask&(1<<(j+1))) == 0 {
									dp[i2][j2][mask2|(1<<(j+1))][k+1] = (dp[i2][j2][mask2|(1<<(j+1))][k+1] + dp[i][j][mask][k]) % MOD
								}
							}
						}
					}
				}
			}
		}
		for i := 0; i < X+1; i++ {
			small[i][Y] = dp[i][0][0][N]
		}
	}
	for i := 1; i <= 2*N; i++ {
		brute(N, i)
	}
	ans := FUNC(N, X, Y)
	fmt.Println(ans)
}
