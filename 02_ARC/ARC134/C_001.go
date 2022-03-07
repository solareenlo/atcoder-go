package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscan(in, &n, &k)

	A := make([]int, n)
	for i := range A {
		fmt.Fscan(in, &A[i])
	}

	s := 0
	for i := 1; i < n; i++ {
		s += A[i]
	}

	if A[0] < s+k {
		fmt.Println(0)
		return
	}

	A[0] -= s + k
	ans := 1
	d := 1
	for i := 0; i < n; i++ {
		a := A[i] + k - 1
		for j := 0; j < k-1; j++ {
			ans *= a - j
			ans %= mod
			d *= j + 1
			d %= mod
		}
	}
	fmt.Println(divMod(ans, d))
}

const mod = 998244353

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
