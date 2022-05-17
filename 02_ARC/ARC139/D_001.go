package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, M, K, X int
	fmt.Fscan(in, &N, &M, &K, &X)

	X--
	c := make([]int, 2001)
	for i := 0; i < N; i++ {
		var A int
		fmt.Fscan(in, &A)
		A--
		c[A]++
	}

	for i := M; i > 0; i-- {
		c[i-1] += c[i]
	}

	a := powMod(M, K) * N % mod

	for i := 1; i < M; i++ {
		t := powMod(i, K)
		for j := 0; j <= K; j++ {
			a += (max(N, X+c[i]+j-K) - max(N-c[i]-j, X)) * t % mod
			a %= mod
			t *= M - i
			t %= mod
			t *= K - j
			t %= mod
			t = divMod(t, i*(j+1))
		}
	}
	fmt.Println(a)
}

const mod = 998244353

func powMod(a, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = res * a % mod
		}
		a = a * a % mod
		n /= 2
	}
	return res
}

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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
