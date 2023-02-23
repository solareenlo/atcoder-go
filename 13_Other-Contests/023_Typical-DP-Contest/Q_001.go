package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const mod = 1000000007

	var m, n int
	fmt.Fscan(in, &m, &n)
	var able [50][1 << 8]int
	for i := 0; i < m; i++ {
		var str string
		fmt.Fscan(in, &str)
		x := 0
		l := len(str)
		for i := 0; i < l; i++ {
			x = x*2 + int(str[i]-'0')
		}
		able[l][x] = 1
	}

	var add func(*int, int)
	add = func(x *int, y int) {
		*x += y
		if *x >= mod {
			*x -= mod
		}
	}

	var dp [2][1 << 7][1 << 8]int
	now := 0
	dp[now][0][1<<7] = 1
	for i := 1; i <= n; i++ {
		now ^= 1
		for j := range dp[now] {
			for k := range dp[now][j] {
				dp[now][j][k] = 0
			}
		}
		for j := 0; j < (1 << 7); j++ {
			for k := 0; k < (1 << 8); k++ {
				if dp[now^1][j][k] != 0 {
					for ui := 0; ui < 2; ui++ {
						ci := 0
						for l := 1; l <= i && ci == 0 && l <= 8; l++ {
							if (k & (1 << (8 - l))) != 0 {
								if able[l][(j|(ui<<7))>>(8-l)] != 0 {
									ci = 1
								}
							}
						}
						add(&dp[now][(ui<<6)|(j>>1)][(ci<<7)|(k>>1)], dp[now^1][j][k])
					}
				}
			}
		}
	}

	ans := 0
	for i := 0; i < (1 << 7); i++ {
		for j := 0; j < (1 << 8); j++ {
			if (j & (1 << 7)) != 0 {
				add(&ans, dp[now][i][j])
			}
		}
	}
	fmt.Println(ans)
}
