package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const MOD = 1000000007
	const MX = 177148

	var n, x, p int
	fmt.Fscan(in, &n, &x, &p)
	vec := make([]int, n)
	for i := 0; i < n; i++ {
		var str string
		fmt.Fscan(in, &str)
		if str == "?" {
			vec[i] = -1
		} else {
			vec[i], _ = strconv.Atoi(str)
		}
	}

	x3 := 3
	for i := 0; i < x; i++ {
		x3 *= 3
	}

	var nmask [11][MX]int
	for mask := 0; mask < x3; mask++ {
		now := make([]int, x+1)
		nm := mask
		c := 0
		for nm > 0 {
			now[c] = nm % 3
			c++
			nm /= 3
		}
		for j := 0; j < p+1; j++ {
			mul := 1
			for k := 0; k < x+1; k++ {
				a := now[k]
				if k-j >= 0 {
					a += now[k-j]
				}
				a = min(a, 2)
				nmask[j][mask] += a * mul
				mul *= 3
			}
		}
	}

	dp := make([][]int, 2)
	for i := range dp {
		dp[i] = make([]int, MX)
	}
	dp[0][1] = 1
	for i := 0; i < n; i++ {
		for j := range dp[1] {
			dp[1][j] = 0
		}
		for mask := 0; mask < x3; mask++ {
			for j := 0; j < p+1; j++ {
				if vec[i] == -1 || vec[i] == j {
					dp[1][nmask[j][mask]] = (dp[1][nmask[j][mask]] + dp[0][mask]) % MOD
				}
			}
		}
		dp[0], dp[1] = dp[1], dp[0]
	}

	ans := 0
	for mask := x3 / 3; mask < x3*2/3; mask++ {
		ans = (ans + dp[0][mask]) % MOD
	}
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
