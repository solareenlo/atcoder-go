package main

import "fmt"

var (
	dx   = []int{0, 0, -1, 0, 1}
	dy   = []int{0, 1, 0, -1, 0}
	memo = make([]float64, 1<<16)
)

func dfs(S int) float64 {
	if memo[S] != -1.0 {
		return memo[S]
	}
	if S == (1<<16)-1 {
		return 0
	}
	memo[S] = 1e9
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			cnt := 0
			tmp := 0.0
			for k := 0; k < 5; k++ {
				x := i + dx[k]
				y := j + dy[k]
				if 0 <= x && x < 4 && 0 <= y && y < 4 {
					S2 := S | 1<<(x*4+y)
					if S2 == S {
						cnt++
					} else {
						tmp += dfs(S2) / 5.0
					}
				} else {
					cnt++
				}
			}
			if cnt < 5 {
				memo[S] = min(memo[S], (tmp+1.0)/(1.0-float64(cnt)/5.0))
			}
		}
	}
	return memo[S]
}

func main() {
	var s string
	for i := 0; i < 4; i++ {
		var t string
		fmt.Scan(&t)
		s += t
	}

	S0 := 0
	for i := 0; i < 16; i++ {
		if s[i] == '.' {
			S0 |= 1 << i
		}
	}

	for S := 0; S < 1<<16; S++ {
		memo[S] = -1.0
	}
	fmt.Println(dfs(S0))
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
