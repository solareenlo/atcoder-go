package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var p [200020]int

	var n, m int
	fmt.Fscan(in, &n, &m)
	ans := 0
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &p[i])
		ans = (ans + i*i%MOD) % MOD
	}
	for i, v, t := 1, powMod(n*(n+1)/2%MOD, MOD-2), (n+1)*(MOD+1)/2%MOD; i <= n; i++ {
		prob := powMod(((i-1)*i+(n-i)*(n-i+1))/2%MOD*v%MOD, m)
		ans = (ans + (prob*i+(MOD+1-prob)*t)%MOD*(MOD-p[i])) % MOD
	}
	fmt.Println(ans * powMod(n*(n+1)/2%MOD, m) % MOD)
}

const MOD = 998244353

func powMod(a, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = res * a % MOD
		}
		a = a * a % MOD
		n /= 2
	}
	return res
}
