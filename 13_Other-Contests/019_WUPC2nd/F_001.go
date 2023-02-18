package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const MAX = 105
	const INF = 1 << 30

	var H, W, L int
	fmt.Fscan(in, &H, &W, &L)
	var T string
	fmt.Fscan(in, &T)
	S := make([]string, H)
	for i := 0; i < H; i++ {
		fmt.Fscan(in, &S[i])
	}

	var dp [MAX][MAX][MAX][4]int
	for i := 1; i < MAX; i++ {
		for j := 0; j < MAX; j++ {
			for k := 0; k < MAX; k++ {
				for d := 0; d < 4; d++ {
					dp[i][j][k][d] = INF
				}
			}
		}
	}

	var dh = [4]int{0, 1, 0, -1}
	var dw = [4]int{1, 0, -1, 0}
	for t := 0; t < L; t++ {
		for i := 0; i < H; i++ {
			for j := 0; j < W; j++ {
				mi := INF
				for k := 0; k < 4; k++ {
					mi = min(mi, dp[t][i][j][k])
				}
				if mi == INF {
					continue
				}

				dis := make([]int, 4)
				nh := make([]int, 4)
				nw := make([]int, 4)

				for k := 0; k < 4; k++ {
					h := i + dh[k]
					w := j + dw[k]
					f := true
					for {
						if h < 0 || h >= H || w < 0 || w >= W {
							f = false
							break
						}
						if S[h][w] == T[t] {
							break
						}
						h += dh[k]
						w += dw[k]
					}
					if f {
						dis[k] = abs(i-h) + abs(j-w)
						nh[k] = h
						nw[k] = w
					} else {
						dis[k] = INF
					}
				}

				for d := 0; d < 4; d++ {
					if dp[t][i][j][d] == INF {
						continue
					}
					for k := 0; k < 4; k++ {
						if d == k {
							continue
						}
						if dis[k] < INF {
							dp[t+1][nh[k]][nw[k]][k] = min(dp[t+1][nh[k]][nw[k]][k], dp[t][i][j][d]+dis[k])
						}
					}
				}
			}
		}
	}

	ans := INF
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			for d := 0; d < 4; d++ {
				ans = min(ans, dp[L][i][j][d])
			}
		}
	}

	if ans == INF {
		ans = -1
	}
	fmt.Println(ans)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
