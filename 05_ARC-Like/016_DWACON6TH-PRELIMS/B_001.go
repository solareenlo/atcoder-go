package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, a int
	fmt.Fscan(in, &n, &a)
	t := 1
	for i := 1; i < n; i++ {
		t *= i
		t %= mod
	}

	k := 0
	ans := 0
	for i := 0; i < n-1; i++ {
		var b int
		fmt.Fscan(in, &b)
		k += divMod(t, (i + 1))
		k %= mod
		ans += ((b - a + mod) % mod) * k % mod
		ans %= mod
		a = b
	}

	fmt.Println(ans)
}

const mod = 1000000007

func divMod(a, b int) int {
	ret := a * modInv(b)
	ret %= mod
	return ret
}

func modInv(a int) int {
	b, u, v := mod, 1, 0
	for b != 0 {
		t := a / b
		a -= t * b
		a, b = b, a
		u -= t * v
		u, v = v, u
	}
	u %= mod
	if u < 0 {
		u += mod
	}
	return u
}
