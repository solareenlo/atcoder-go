package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	C := [62][62]int{}
	m := n<<1 | 1
	for i := 0; i <= m; i++ {
		C[i][0] = 1
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= i; j++ {
			C[i][j] = C[i-1][j] + C[i-1][j-1]
		}
	}

	dp := [62][62][2002]int{}
	for i := 1; i <= m; i++ {
		if i*(i-1)/2 <= k {
			dp[i][i-1][i*(i-1)/2] = 1
		}
	}

	for i := 1; i <= m; i++ {
		for j := 0; j < i; j++ {
			for t := 0; t <= k; t++ {
				if dp[i][j][t] != 0 {
					for x := j + 2; i+x <= m && t+x*(x-1)/2 <= k; x++ {
						dp[i+x][x-j-2][t+x*(x-1)/2] += dp[i][j][t] * C[x-1][j+1]
					}
				}
			}
		}
	}

	s := dp[m][0][k]
	for i := 1; i < m; i++ {
		for j := 1; j <= m; j++ {
			for t := 0; t <= k; t++ {
				if dp[i][j][t] != 0 {
					s += dp[i][j][t] * dp[m-i][j-1][k-t]
				}
			}
		}
	}
	fmt.Println(s)
}
