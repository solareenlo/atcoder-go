package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

const INF = int(1e18)

var H, W int
var C, D, S, T [16]int
var s [100]string
var dist [100][100]int
var d [5]int = [5]int{0, 1, 0, -1, 0}

func main() {
	in := bufio.NewReader(os.Stdin)

	var G [16][16]int
	var dp [1 << 16][16]int
	var N int
	fmt.Fscan(in, &H, &W, &N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &C[i], &D[i], &S[i], &T[i])
		C[i]--
		D[i]--
	}
	for i := 0; i < H; i++ {
		fmt.Fscan(in, &s[i])
	}
	for i := 0; i < N; i++ {
		mkd(C[i], D[i])
		for j := 0; j < N; j++ {
			G[i][j] = dist[C[j]][D[j]]
		}
	}
	for i := 0; i < 1<<N; i++ {
		for j := 0; j < N; j++ {
			dp[i][j] = INF
		}
	}
	mkd(0, 0)
	for i := 0; i < N; i++ {
		nd := dist[C[i]][D[i]]
		if nd == INF {
			continue
		}
		if nd < S[i] {
			if nd%2 != S[i]%2 {
				nd = S[i] + 1
			} else {
				nd = S[i]
			}
		}
		if S[i] <= nd && nd <= T[i] {
			dp[1<<i][i] = nd
		}
	}
	ans := 0
	for i := 1; i < 1<<N; i++ {
		for j := 0; j < N; j++ {
			if dp[i][j] < INF {
				ans = max(ans, popcount(uint32(i)))
				for k := 0; k < N; k++ {
					if ((i>>k)&1) == 0 && G[j][k] < INF {
						nd := dp[i][j] + G[j][k]
						if nd < S[k] {
							if nd%2 != S[k]%2 {
								nd = S[k] + 1
							} else {
								nd = S[k]
							}
						}
						if S[k] <= nd && nd <= T[k] {
							dp[i|(1<<k)][k] = min(dp[i|(1<<k)][k], nd)
						}
					}
				}
			}
		}
	}
	fmt.Println(ans)
}

func mkd(sx, sy int) {
	type pair struct {
		x, y int
	}

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			dist[i][j] = INF
		}
	}
	dist[sx][sy] = 0
	P := make([]pair, 0)
	P = append(P, pair{sx, sy})
	for len(P) > 0 {
		x := P[0].x
		y := P[0].y
		P = P[1:]
		for r := 0; r < 4; r++ {
			tx := x + d[r]
			ty := y + d[r+1]
			if tx < 0 || ty < 0 || tx >= H || ty >= W || s[tx][ty] == '#' || dist[tx][ty] <= dist[x][y]+1 {
				continue
			}
			dist[tx][ty] = dist[x][y] + 1
			P = append(P, pair{tx, ty})
		}
	}
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
