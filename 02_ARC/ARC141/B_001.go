package main

import "fmt"

func main() {
	const mod = 998244353

	var m, lim int
	fmt.Scan(&m, &lim)

	if m >= 80 {
		fmt.Println(0)
		return
	}

	t := 0
	for (1 << t) <= lim {
		t++
	}

	var f [100][100]int
	f[0][0] = 1
	for i := 1; i <= t; i++ {
		now := 1 << (i - 1)
		if i == t {
			now = lim - (1 << (t - 1)) + 1
		}
		now %= mod
		for j := 1; j <= t; j++ {
			for x := 0; x < i; x++ {
				f[i][j] = (f[i][j] + f[x][j-1]*now) % mod
			}
		}
	}

	ans := 0
	for i := 1; i <= t; i++ {
		ans = (ans + f[i][m]) % mod
	}
	fmt.Println(ans)
}
