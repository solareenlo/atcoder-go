package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const MOD = 998244353
	const N = 200005

	var a [N]int
	var vis [N]bool

	var n, m int
	fmt.Fscan(in, &n, &m)
	ans := 1
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		ans = ans * m % MOD
	}
	s := 1
	for i := 1; i <= n; i++ {
		if vis[i] {
			continue
		}
		s = s * m % MOD
		u := i
		for !vis[u] {
			vis[u] = true
			u = a[u]
		}
	}
	ans = (ans - s + MOD) % MOD * 499122177 % MOD
	fmt.Println(ans)
}
