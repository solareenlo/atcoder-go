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

	y := make([]int, n+1)
	for i := range y {
		fmt.Fscan(in, &y[i])
	}

	var t int
	fmt.Fscan(in, &t)

	P := (t - n + mod) % mod
	a := 1
	for i := 0; i < n; i++ {
		P *= (t - i + mod) % mod
		P %= mod
		a *= (-i - 1 + mod) % mod
		a %= mod
	}

	res := 0
	for i := 0; i < n+1; i++ {
		tmp := divMod(y[i]*P%mod, (t-i+mod)%mod)
		res += divMod(tmp, a)
		res %= mod
		if i < n {
			a *= i + 1
			a %= mod
			a = divMod(a, (i-n+mod)%mod)
		}
	}
	fmt.Println(res)
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
