package main

import "fmt"

const mod = 998244353

func F(n, k int) int {
	for i := 0; i < k; i++ {
		n = (n*3 + 2) / 2
	}
	return n
}

func FF(n int) int {
	n %= mod
	return n * (n - 1) / 2 % mod
}

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	m := n
	for i := 1; i <= k; i++ {
		m -= (m + 2) / 3
	}
	if k > 38 {
		ans := 0
		for i := 0; i < m; i++ {
			ans = (ans + F(i, k) + 1) % mod
		}
		fmt.Println(ans)
		return
	}

	if k <= 20 {
		t := 1
		for i := 0; i < k; i++ {
			t *= 3
		}
		ans := m % mod
		t %= mod
		for i := 0; i < 1<<k && i < m; i++ {
			c := ((m - i - 1) >> k) + 1
			v := F(i, k)
			ans = (ans + FF(c)*t + v*c) % mod
		}
		fmt.Println(ans)
		return
	}

	X := k / 2
	Y := k - X
	t := 1
	for i := 0; i < X; i++ {
		t *= 3
	}
	tt := 1
	for i := 0; i < Y; i++ {
		tt = 3 * tt % mod
	}
	L := 0
	for (1 << L) <= (m >> X) {
		L++
	}
	dp := [26][1 << 19]int{}
	for i := 0; i < 1<<Y; i++ {
		dp[0][i] = F(i, Y) % mod
	}
	for j := 1; j <= L; j++ {
		for i := 0; i < 1<<Y; i++ {
			to := i + (1<<(j-1))*t
			dp[j][i] = (dp[j-1][i] + dp[j-1][to%(1<<Y)] + (to>>Y)%mod*(1<<(j-1))%mod*tt) % mod
		}
	}

	ans := m % mod
	for i := 0; i < 1<<X && i < m; i++ {
		c := ((m - i - 1) >> X) + 1
		tot := 0
		v := F(i, X)
		tot = (v >> Y) * tt % mod
		v %= 1 << Y
		for j := 0; j <= L; j++ {
			if c>>j&1 != 0 {
				ans = (ans + (tot << j) + dp[j][v]) % mod
				v += (t << j)
				tot = (tot + (v>>Y)*tt) % mod
				v %= 1 << Y
			}
		}
	}
	fmt.Println(ans)
}
