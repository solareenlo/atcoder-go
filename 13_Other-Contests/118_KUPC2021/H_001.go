package main

import (
	"bufio"
	"fmt"
	"os"
)

const MOD = 998244353

var s, t, u int

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m, &s, &t, &u)

	ret := make([][]int, 60)
	for i := range ret {
		ret[i] = make([]int, 3)
	}
	ret[0][1] = 1

	for i := 1; i < 60; i++ {
		ret[i] = f(ret[i-1], ret[i-1])
	}

	A := []int{1, 0, 0}
	B := []int{1, 0, 0}
	for i := 0; i < 60; i++ {
		if n%2 == 1 {
			A = f(A, ret[i])
		}
		if m%2 == 1 {
			B = f(B, ret[i])
		}
		n >>= 1
		m >>= 1
	}

	ans := A[1]*B[2]%MOD + MOD - A[2]*B[1]%MOD
	fmt.Println(ans % MOD)
}

func f(a, b []int) []int {
	ret := make([]int, 5)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			ret[i+j] += a[i] * b[j]
			ret[i+j] %= MOD
		}
	}
	ret[3] = (ret[3] + ret[4]*s) % MOD
	ret[2] = (ret[2] + ret[4]*(MOD-t)) % MOD
	ret[1] = (ret[1] + ret[4]*u) % MOD
	ret[4] = 0
	ret[2] = (ret[2] + ret[3]*s) % MOD
	ret[1] = (ret[1] + ret[3]*(MOD-t)) % MOD
	ret[0] = (ret[0] + ret[3]*u) % MOD
	ret[3] = 0
	resize(&ret, 3)
	return ret
}

func resize(a *[]int, n int) {
	if len(*a) > n {
		*a = (*a)[:n]
	} else {
		n = n - len(*a)
		for i := 0; i < n; i++ {
			*a = append(*a, 0)
		}
	}
}
