package main

import (
	"bufio"
	"fmt"
	"os"
)

const mod = 998244353
const N = 2005

var fa, siz, dp, f [N]int

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	for i := 1; i <= n; i++ {
		fa[i] = i
		siz[i] = 1
	}

	sum := 1
	for i := 1; i <= n; i++ {
		var x int
		fmt.Fscan(in, &x)
		if x < 0 {
			sum = sum * n % mod
		} else {
			merge(x, i)
		}
	}

	ans := 0
	dp[0] = 1
	for i := 1; i <= n; i++ {
		if find(i) == i {
			if f[i] != 0 {
				ans = (ans + sum) % mod
				continue
			}
			for j := n; j >= 0; j-- {
				dp[j] = dp[j] * n % mod
				if j != 0 {
					dp[j] = (dp[j] + dp[j-1]*siz[i]) % mod
				}
			}
		}
	}

	num := 1
	for i := 1; i <= n; i++ {
		ans = (ans + dp[i]*num) % mod
		num = num * i % mod
	}
	fmt.Println(ans)
}

func find(x int) int {
	if x == fa[x] {
		return x
	}
	fa[x] = find(fa[x])
	return fa[x]
}

func merge(x, y int) {
	x = find(x)
	y = find(y)
	if x == y {
		f[x] = 1
		return
	}
	fa[x] = y
	siz[y] += siz[x]
	f[y] |= f[x]
}
