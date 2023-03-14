package main

import (
	"bufio"
	"fmt"
	"os"
)

const mod = 1000000007

var fact, inv, factinv, factfact [1100000]int

func main() {
	in := bufio.NewReader(os.Stdin)

	fact[0] = 1
	for i := 1; i < 1100000; i++ {
		fact[i] = fact[i-1] * i % mod
	}
	inv[1] = 1
	for i := 2; i < 1100000; i++ {
		inv[i] = (mod - (mod/i)*inv[mod%i]%mod) % mod
	}
	factinv[0] = 1
	for i := 1; i < 1100000; i++ {
		factinv[i] = factinv[i-1] * inv[i] % mod
	}
	factfact[0] = 1
	for i := 1; i < 1100000; i++ {
		factfact[i] = factfact[i-1] * fact[i] % mod
	}

	var T int
	fmt.Fscan(in, &T)
	for i := 0; i < T; i++ {
		var a, b, c, d int
		fmt.Fscan(in, &a, &b, &c, &d)
		a--
		b--
		c--
		d--
		A := a - b
		B := b
		C := c - d
		D := d
		if A > C || (A == C && B > D) {
			A, C = C, A
			B, D = D, B
		}
		if A <= C && B <= D {
			N := (C + 1) * (D + 1)
			ret := fact[N]
			ret = ret * getinv(factfact[C+D+1]) % mod * factfact[C] % mod
			ret = ret * factfact[D] % mod
			fmt.Println(ret)
		} else {
			N := (C+1)*(B+1) - (C-A)*(B-D)
			ret := fact[N]
			E := C - A
			F := B - D
			ret = ret * getinv(factfact[A+F]) % mod * factfact[F-1] % mod
			ret = ret * factfact[A] % mod
			ret = ret * getinv(factfact[D+E]) % mod * factfact[E-1] % mod
			ret = ret * factfact[D] % mod
			ret = ret * getinv(factfact[B+C+1]) % mod * factfact[B+C-A] % mod
			ret = ret * factfact[B+C-D] % mod * getinv(factfact[B+C-A-D-1]) % mod
			fmt.Println(ret)
		}
	}
}

func getinv(a int) int {
	ret := 1
	pw := mod - 2
	for pw > 0 {
		if pw%2 != 0 {
			ret = ret * a % mod
		}
		pw /= 2
		a = a * a % mod
	}
	return ret
}
