package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const mod = int(1e9) + 7

	var n, k int
	fmt.Fscan(in, &n, &k)
	var s string
	fmt.Fscan(in, &s)
	l := len(s)
	k -= l
	if k < 0 {
		fmt.Println(0)
		return
	}

	ans := 0
	LEN := 0
	for i := 0; i < int(s[0]-'a'); i++ {
		ok := true
		for j := 0; j < l; j++ {
			if int(s[j]) == ('a' + i) {
				ok = false
			}
		}
		if ok {
			LEN++
		}
	}
	ll := int(s[0] - 'a')
	sum := n*ll - k
	if sum < 0 {
		fmt.Println(0)
		return
	}
	f := make([]int, n+1)
	f[0] = 1
	for i := 1; i <= n; i++ {
		f[i] = (f[i-1] * i) % mod
	}
	dp := make([][][]int, n)
	for i := range dp {
		dp[i] = make([][]int, LEN+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, sum+1)
		}
	}
	for i := l - 1; i < n; i++ {
		ss := 0
		for j := 1; j < l; j++ {
			if s[j] < s[0] {
				ss += i + j - l + 1
			}
		}
		sum -= ss
		if sum >= 0 {
			for j := 0; j < n; j++ {
				for u := 0; u <= LEN; u++ {
					for sm := 0; sm <= sum; sm++ {
						if j != 0 {
							dp[j][u][sm] = dp[j-1][u][sm]
							if u != 0 && sm >= j && (j < i-l+1 || j > i) {
								dp[j][u][sm] = (dp[j][u][sm] + dp[j-1][u-1][sm-j]) % mod
							}
						} else {
							dp[j][u][sm] = 0
							if u == 0 && sm == 0 {
								dp[j][u][sm]++
							} else if u == 1 && sm == 0 && i-l+1 > 0 {
								dp[j][u][sm]++
							}
						}
					}
				}
			}
			x := dp[n-1][LEN][sum]
			x = (x * f[LEN]) % mod
			x = (x * f[n-l-LEN]) % mod
			ans = (ans + x) % mod
		}
		sum += ss
	}
	fmt.Println(ans)
}
