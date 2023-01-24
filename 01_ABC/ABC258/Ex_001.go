package main

import (
	"bufio"
	"fmt"
	"os"
)

const MOD = 998244353

type A [4]int

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, s int
	fmt.Fscan(in, &n, &s)

	last := 0
	ans := A{1, 0, 0, 1}
	var a int
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a)
		ans = mul(ans, pow(A{0, 1, 1, 1}, a-last-1))
		ans = mul(ans, A{0, 1, 1, 0})
		last = a
	}
	ans = mul(ans, pow(A{0, 1, 1, 1}, s-last))
	fmt.Println(ans[2])
}

func mul(a, b A) A {
	return A{(a[0]*b[0] + a[1]*b[2]) % MOD,
		(a[0]*b[1] + a[1]*b[3]) % MOD,
		(a[2]*b[0] + a[3]*b[2]) % MOD,
		(a[2]*b[1] + a[3]*b[3]) % MOD}
}

func pow(a A, k int) A {
	res := A{1, 0, 0, 1}
	r := a
	for k > 0 {
		if k&1 != 0 {
			res = mul(res, r)
		}
		r = mul(r, r)
		k >>= 1
	}
	return res
}
