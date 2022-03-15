package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 5005
const mod = 1000000007

var (
	to  = [N * 2]int{}
	nx  = [N * 2]int{}
	he  = [N]int{}
	p   = [N]int{}
	o   int
	u   int
	vis = make([]bool, N)
)

func wk(w, x, y int) {
	to[w] = y
	nx[w] = he[x]
	he[x] = w
}

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

func dfs(x int) {
	if p[x] < 0 {
		o = 1
	}
	u++
	vis[x] = true
	for i := he[x]; i > 0; i = nx[i] {
		y := to[i]
		if !vis[y] {
			dfs(y)
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &p[i])
		if ^p[i] != 0 {
			wk(i*2, i, p[i])
			wk(i*2+1, p[i], i)
		}
	}

	t := 0
	ans := 0
	s := make([]int, N)
	for i := 1; i <= n; i++ {
		if !vis[i] {
			u = 0
			o = 0
			dfs(i)
			if o != 0 {
				t++
				s[t] = u
			} else {
				ans++
			}
		}
	}

	ans = powMod(n-1, t) * ans % mod
	dp := make([]int, N)
	dp[0] = 1
	for i := 1; i <= t; i++ {
		for j := i; j > 0; j-- {
			dp[j] = (dp[j-1]*s[i] + dp[j]) % mod
		}
	}
	dp[1] -= t
	for i, c := 1, 1; i <= t; i++ {
		ans = (dp[i]*c%mod*powMod(n-1, t-i) + ans) % mod
		c = c * i % mod
	}
	fmt.Println((n*powMod(n-1, t) + mod - ans) % mod)
}
