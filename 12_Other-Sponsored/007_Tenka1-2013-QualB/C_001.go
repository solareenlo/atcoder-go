package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const INF = int(1e18)

	var dp [2][1 << 8][37][37][2]int

	var N, M, A, B int
	fmt.Fscan(in, &N, &M, &A, &B)
	for i := range dp {
		for j := range dp[i] {
			for k := range dp[i][j] {
				for l := range dp[i][j][k] {
					dp[i][j][k][l][0] = -INF
					dp[i][j][k][l][1] = -INF
				}
			}
		}
	}
	dp[1][0][0][0][0] = 0
	dp[1][0][0][0][1] = 0
	dp[1][1][0][0][0] = 0
	dp[1][1][0][0][1] = 0
	now := 1
	for i := 1; i < (N-1)*M; i++ {
		for j := 0; j < (1 << M); j++ {
			for k := 0; k < A+1; k++ {
				for l := 0; l <= B; l++ {
					dp[1-now][j][k][l][0] = -INF
					dp[1-now][j][k][l][1] = -INF
				}
			}
		}
		for j := 0; j < (1 << M); j++ {
			for k := 0; k < A+1; k++ {
				for l := 0; l <= B; l++ {
					for m := 0; m < 2; m++ {
						t := dp[now][j][k][l][m]
						if t < 0 {
							continue
						}
						x := i / M
						y := i % M
						for r := 0; r < 4; r++ {
							if y == M-1 && r >= 2 {
								break
							}
							nj := j & ^(1<<y) | ((r & 1) << y)
							nm := (r >> 1) & 1
							if 0 < x && x < N-1 && 0 < y && y < M-1 {
								a, b := 0, 0
								tmp0 := 0
								if ((j >> y) & 1) != 0 {
									tmp0 = 1
								}
								tmp1 := 0
								if (r & 1) == 0 {
									tmp1 = 1
								}
								tmp2 := 0
								if nm == 0 {
									tmp2 = 1
								}
								if m == tmp0 && m == tmp1 && m == tmp2 {
									if m == 1 {
										a = 1
									} else {
										b = 1
									}
								}
								if l+b > B {
									continue
								}
								nk := k + a
								if k == A {
									nk = A
								}
								dp[1-now][nj][nk][l+b][nm] = max(dp[1-now][nj][nk][l+b][nm], t)
							} else if x == 0 && y == M-1 {
								dp[1-now][nj][k][l][0] = max(dp[1-now][nj][k][l][0], t)
							} else {
								var tmp bool
								if x == 0 {
									tmp = m != 0 && r == 0
								} else if y == 0 {
									tmp = (j&1) != 0 && r == 0
								} else {
									tmp = (j>>(M-1)) != 0 && m != 0 && r == 0
								}
								if tmp {
									dp[1-now][nj][k][l][nm] = max(dp[1-now][nj][k][l][nm], t+1)
								} else {
									dp[1-now][nj][k][l][nm] = max(dp[1-now][nj][k][l][nm], t)
								}
							}
						}
					}
				}
			}
		}
		now = 1 - now
	}
	ans := 0
	for j := 0; j < (1 << M); j++ {
		for l := 0; l <= B; l++ {
			t := dp[now][j][A][l][0]
			if t < 0 {
				continue
			}
			for i := 0; i < 1<<(M-1); i++ {
				x := t
				for k := 1; k < M-1; k++ {
					if ((j>>k)&1) != 0 && ((i>>(k-1))&1) != 0 && ((i>>k)&1) == 0 {
						x++
					}
				}
				ans = max(ans, x)
			}
		}
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
