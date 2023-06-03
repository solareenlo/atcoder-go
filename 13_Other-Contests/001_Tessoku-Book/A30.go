package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, r int
	fmt.Fscan(in, &n, &r)
	ans := 1
	div := 1
	for i := 0; i < r; i++ {
		ans = ans * (n - i) % mod
		div = div * (i + 1) % mod
	}
	fmt.Println(divMod(ans, div))
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
