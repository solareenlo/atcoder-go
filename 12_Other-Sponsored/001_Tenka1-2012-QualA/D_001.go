package main

import (
	"bufio"
	"fmt"
	"os"
)

var N int
var dx [4]int = [4]int{0, 1, 0, -1}
var dy [4]int = [4]int{1, 0, -1, 0}
var S [51]string
var C, G [51][51]int
var dp [2600][2600]float64

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &S[i])
	}
	dfs(0, 0, 0, 0)
	A := dfs2(0, 0, -1, 0)
	ret := 0.0
	for i := 1; i <= len(A); i++ {
		if i != len(A) {
			ret += (A[i] - A[i-1]) * float64(i)
		}
	}
	fmt.Println(ret)
}

func dfs(x, y, px, py int) {
	C[y][x] = 1
	if x == N-1 && y == N-1 {
		G[y][x] = 1
		return
	}
	for i := 0; i < 4; i++ {
		ty := y + dy[i]
		tx := x + dx[i]
		if ty < 0 || ty >= N || tx < 0 || tx >= N {
			continue
		}
		if ty == py && tx == px {
			continue
		}
		if S[ty][tx] == '#' {
			continue
		}
		dfs(tx, ty, x, y)
		C[y][x] += C[ty][tx]
		G[y][x] += G[ty][tx]
	}
}

func dfs2(x, y, px, py int) []float64 {
	if x == N-1 && y == N-1 {
		return []float64{0, 1.0}
	}
	ox := x + (x - px)
	oy := y + (y - py)
	add := 0
	if ox >= 0 && oy >= 0 && ox < N && oy < N && S[oy][ox] != '#' {
		if G[oy][ox] != 0 {
			return dfs2(ox, oy, x, y)
		}
		add = C[oy][ox]
	}
	tmp := make([]float64, 0)
	num := 0
	for i := 0; i < 4; i++ {
		ty := y + dy[i]
		tx := x + dx[i]
		if ty < 0 || ty >= N || tx < 0 || tx >= N {
			continue
		}
		if ty == py && tx == px {
			continue
		}
		if ty == oy && tx == ox {
			continue
		}
		if S[ty][tx] == '#' {
			continue
		}
		if G[ty][tx] != 0 {
			tmp = dfs2(tx, ty, x, y)
		} else {
			num = C[ty][tx]
		}
	}
	ret := make([]float64, len(tmp)+num+add)
	for a := 0; a < len(tmp); a++ {
		for b := 0; b <= num; b++ {
			dp[a][b] = 0
		}
	}
	dp[0][0] = 1
	for a := 0; a < len(tmp); a++ {
		for b := 0; b <= num; b++ {
			if dp[a][b] != 0 {
				if a+1 < len(tmp) && b+1 <= num {
					dp[a+1][b] += dp[a][b] / 2
					dp[a][b+1] += dp[a][b] / 2
				} else {
					dp[a+1][b] += dp[a][b]
					dp[a][b+1] += dp[a][b]
				}
				ret[add+a+b] += dp[a][b] * tmp[a]
			}
		}
	}
	return ret
}
