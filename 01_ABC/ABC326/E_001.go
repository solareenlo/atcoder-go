package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	ans := 0
	x := 1
	y := n
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(in, &a)
		ans = (ans + divMod((a*x%MOD), y)) % MOD
		x = x * (n + 1) % MOD
		y = y * n % MOD
	}
	fmt.Println(ans)
}

const MOD = 998244353

func divMod(a, b int) int {
	ret := a * modInv(b)
	ret %= MOD
	return ret
}

func modInv(a int) int {
	b, u, v := MOD, 1, 0
	for b != 0 {
		t := a / b
		a -= t * b
		a, b = b, a
		u -= t * v
		u, v = v, u
	}
	u %= MOD
	if u < 0 {
		u += MOD
	}
	return u
}
