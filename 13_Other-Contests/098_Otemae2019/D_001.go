package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const MOD = 1000000007

	var n, m int
	fmt.Fscan(in, &n, &m)

	var c [1005]int
	for i := 0; i < m; i++ {
		var s string
		fmt.Fscan(in, &s)
		if s == "FizzBuzz" {
			c[i] = 3
		} else if s == "Fizz" {
			c[i] = 2
		} else if s == "Buzz" {
			c[i] = 1
		}
	}

	var dp [1005][1005][3]int
	dp[0][0][0] = 1
	for i := 0; i < n; i++ {
		for k := 0; k < 10; k++ {
			if i != 0 || k != 0 {
				for j := 0; j < i+1; j++ {
					for l := 0; l < 3; l++ {
						tmp0 := 0
						if k%5 == 0 {
							tmp0 = 1
						}
						tmp1 := 0
						if (k+l)%3 == 0 {
							tmp1 = 1
						}
						t := tmp0 + tmp1*2
						if t == 0 {
							dp[i+1][j][(k+l)%3] = (dp[i+1][j][(k+l)%3] + dp[i][j][l]) % MOD
						} else if t == c[j] {
							dp[i+1][j+1][(k+l)%3] = (dp[i+1][j+1][(k+l)%3] + dp[i][j][l]) % MOD
						}
					}
				}
			}
		}
	}

	ans := 0
	for i := 0; i < 3; i++ {
		ans = (ans + dp[n][m][i]) % MOD
	}
	fmt.Println(ans)
}
