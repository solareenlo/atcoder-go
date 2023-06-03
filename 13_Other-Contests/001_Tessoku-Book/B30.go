package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var h, w int
	fmt.Fscan(in, &h, &w)
	h--
	w--
	r := min(h, w)
	ans := 1
	div := 1
	h += w
	for i := 0; i < r; {
		ans = ans * h % mod
		h--
		i++
		div = div * i % mod
	}
	fmt.Println(divMod(ans, div))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
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
