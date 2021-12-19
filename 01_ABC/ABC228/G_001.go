package main

import "fmt"

func main() {
	var H, W, N int
	fmt.Scan(&H, &W, &N)

	S := make([]string, 10)
	for i := 0; i < H; i++ {
		fmt.Scan(&S[i])
	}

	type pair struct{ x, y int }
	to := make([][]pair, 100)
	for j := 0; j < W; j++ {
		for i := 0; i < H; i++ {
			for k := 0; k < H; k++ {
				to[(S[i][j]-'0')*10+(S[k][j]-'0')] = append(to[(S[i][j]-'0')*10+(S[k][j]-'0')], pair{i, k})
			}
		}
	}

	G := make([][]int, 1<<10)
	for i := 0; i < 100; i++ {
		for x := 1; x < 1<<H; x++ {
			y := 0
			for _, e := range to[i] {
				if x>>e.x&1 != 0 {
					y |= 1 << e.y
				}
			}
			if y > 0 {
				G[x] = append(G[x], y)
			}
		}
	}

	mod := 998244353
	dp := [301][1 << 10]int{}
	dp[0][(1<<H)-1] = 1
	for t := 0; t < N; t++ {
		for x := 1; x < 1<<H; x++ {
			for _, y := range G[x] {
				dp[t+1][y] += dp[t][x]
				dp[t+1][y] %= mod
			}
		}
	}

	res := 0
	for x := 1; x < 1<<H; x++ {
		res += dp[N][x]
		res %= mod
	}
	fmt.Println(res)
}
