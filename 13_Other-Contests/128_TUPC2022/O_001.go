package main

import (
	"bufio"
	"fmt"
	"os"
)

var C [30]int
var first int

func main() {
	in := bufio.NewReader(os.Stdin)

	const MOD = 998244353

	var n, m int
	var T string
	fmt.Fscan(in, &n, &m, &T)
	first = n
	now := 0
	for i := 0; i < n; i++ {
		for T[i+now] == '1' {
			now += 1
			if first == n {
				first = i
			}
		}
		C[i] = now
	}

	offset := n*m + 5
	var dp [32][2 * (32*32 + 5)][32][32]int
	dp[0][offset][0][0] = 1
	for i := 0; i < n; i++ {
		var cum [2 * (32*32 + 5)][32][32]int
		for j := -i * m; j <= i*m; j++ {
			for a := 0; a <= m; a++ {
				for b := 0; b <= m; b++ {
					cum[j+a-b+offset][a][b] = dp[i][j+offset][a][b]
					if a > 0 {
						cum[j+a-b+offset][a][b] = (cum[j+a-b+offset][a][b] + cum[j+(a-1)-b+offset][a-1][b]) % MOD
					}
					if b > 0 {
						cum[j+a-b+offset][a][b] = (cum[j+a-b+offset][a][b] + cum[j+a-(b-1)+offset][a][b-1]) % MOD
					}
					if a > 0 && b > 0 {
						cum[j+a-b+offset][a][b] = (cum[j+a-b+offset][a][b] - cum[j+(a-1)-(b-1)+offset][a-1][b-1] + MOD) % MOD
					}
				}
			}
		}

		for a := 0; a <= m; a++ {
			for b := 0; b <= m; b++ {
				if check(i, a, b) {
					for j := -(i + 1) * m; j <= (i+1)*m; j++ {
						if j-2*max(C[i]-b, 0)+offset < 0 {
							continue
						}
						dp[i+1][j-2*max(C[i]-b, 0)+offset][a][b] = (dp[i+1][j-2*max(C[i]-b, 0)+offset][a][b] + cum[j+offset][a][b]) % MOD
					}
				}
			}
		}
	}

	ans := 0
	for a := 0; a <= m; a++ {
		for b := 0; b <= m; b++ {
			ans = (ans + dp[n][offset][a][b]) % MOD
		}
	}
	if first != n {
		ans = ans * 2 % MOD
	}
	fmt.Println(ans)
}

func check(i, a, b int) bool {
	if i < first {
		return true
	} else if i == first {
		return (b < C[i] && C[i] <= a)
	} else if C[i-1] < C[i] {
		return (b < C[i] && C[i] == a)
	}
	return max(b, C[i]) >= a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
