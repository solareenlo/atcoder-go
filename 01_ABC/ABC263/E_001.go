package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 202020

	var n int
	fmt.Fscan(in, &n)
	var a [N]int
	for i := 1; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	var s, dp [N]int
	for i := n - 1; i > 0; i-- {
		dp[i] = (s[i+1] - s[i+a[i]+1] + a[i] + 1 + mod) * powMod(a[i], mod-2) % mod
		s[i] = (s[i+1] + dp[i]) % mod
	}
	fmt.Println((dp[1] + mod) % mod)
}

const mod = 998244353

func powMod(a, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = res * a % mod
		}
		a = a * a % mod
		n /= 2
	}
	return res
}
