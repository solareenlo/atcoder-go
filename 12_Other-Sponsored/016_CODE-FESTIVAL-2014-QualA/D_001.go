package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const INF = int(1e18)

	var A string
	var K int
	fmt.Fscan(in, &A, &K)

	var dp [17][1 << 10][3]int

	for i := 0; i <= len(A); i++ {
		for j := 0; j < 1<<10; j++ {
			dp[i][j][0] = INF
			dp[i][j][1] = INF
			dp[i][j][2] = INF
		}
	}
	dp[0][0][0] = 0
	for i := 0; i < len(A); i++ {
		d := int(A[i] - '0')
		for j := 0; j < 1<<10; j++ {
			for k := 0; k < 3; k++ {
				if dp[i][j][k] > int(1e17) {
					continue
				}
				for l := 0; l < 10; l++ {
					nj := j | (1 << l)
					if k == 0 {
						if d == l {
							dp[i+1][nj][0] = 0
						} else if d < l {
							dp[i+1][nj][1] = min(dp[i+1][nj][1], l-d)
						} else {
							dp[i+1][nj][2] = min(dp[i+1][nj][2], d-l)
						}
					} else if k == 1 {
						dp[i+1][nj][1] = min(dp[i+1][nj][1], dp[i][j][1]*10+(l-d))
					} else {
						dp[i+1][nj][2] = min(dp[i+1][nj][2], dp[i][j][2]*10+(d-l))
					}
				}
			}
		}
	}

	ans := INF
	for j := 0; j < 1<<10; j++ {
		if popcount(uint32(j)) <= K {
			for k := 0; k < 3; k++ {
				ans = min(ans, dp[len(A)][j][k])
			}
		}
	}
	a := 0
	for i := 0; i < len(A); i++ {
		a = a*10 + int(A[i]-'0')
	}
	t := 1
	for i := 1; i < len(A); i++ {
		t *= 10
	}
	ans = min(ans, a-(t-1))

	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func popcount(n uint32) int {
	return bits.OnesCount32(n)
}
