package main

import "fmt"

const MOD = 998244353
const MAXD = 1004

func powMod(a, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = res * a % MOD
		}
		a = a * a % MOD
		n /= 2
	}
	return res
}

var (
	n    int
	a    = [1005]int{}
	rem  = [1005]int{}
	fac  = [1005]int{}
	ifac = [1005]int{}
	inv  = [1005]int{}
	S0   = [1005][1005]int{}
	S1   = [1005][1005]int{}
	M    int
	p    = [1005]int{}
	q    = [1005]int{}
	r    = [1005]int{}
)

func mul1_p() {
	for i := n; i >= 0; i-- {
		p[i+1] = (p[i+1] + p[i]) % MOD
	}
}

func convm_p(x, d int) {
	for i := 0; i < n+2; i++ {
		q[i] = 0
		r[i] = 0
	}
	for i := 0; i < n+2; i++ {
		for j := 0; j < i+1; j++ {
			q[j] = (q[j] + p[i]*S0[i][j]) % MOD
		}
	}
	for i := n + 1; i >= 0; i-- {
		for j := n + 1 - i; j >= 0; j-- {
			tmp := 0
			if j > 0 {
				tmp = x * r[j-1]
			}
			r[j] = (d*r[j] + tmp) % MOD
		}
		r[0] = (r[0] + q[i]) % MOD
	}
	for i := 0; i < n+2; i++ {
		p[i] = 0
	}
	for i := 0; i < n+2; i++ {
		for j := 0; j < i+1; j++ {
			p[j] = (p[j] + r[i]*S1[i][j]) % MOD
		}
	}
}

func calc_p(x int) int {
	coef := 1
	ret := 0
	for i := 0; i < n+2; i++ {
		ret = (ret + p[i]*coef) % MOD
		coef = coef * (x - i + MOD) % MOD * inv[i+1] % MOD
	}
	return ret
}

func main() {
	fac[0] = 1
	for i := 1; i < MAXD+1; i++ {
		fac[i] = fac[i-1] * i % MOD
	}
	ifac[MAXD] = powMod(fac[MAXD], MOD-2)
	for i := MAXD; i >= 1; i-- {
		ifac[i-1] = ifac[i] * i % MOD
	}
	for i := 1; i < MAXD+1; i++ {
		inv[i] = ifac[i] * fac[i-1] % MOD
	}
	S0[0][0] = 1
	for i := 1; i < MAXD+1; i++ {
		for j := 1; j < i+1; j++ {
			S0[i][j] = ((MOD-i+1)*S0[i-1][j] + S0[i-1][j-1]) % MOD
		}
	}
	for i := 1; i < MAXD+1; i++ {
		for j := 1; j < i+1; j++ {
			S0[i][j] = S0[i][j] * ifac[i] % MOD
		}
	}
	S1[0][0] = 1
	for i := 1; i < MAXD+1; i++ {
		for j := 1; j < i+1; j++ {
			S1[i][j] = j * (S1[i-1][j-1] + S1[i-1][j]) % MOD
		}
	}

	fmt.Scan(&n, &M)
	for i := 1; i < n+1; i++ {
		fmt.Scan(&a[i])
	}

	for i := n; i >= 1; i-- {
		rem[i] = M % a[i]
		M /= a[i]
	}
	rem[0] = (M - 1) % MOD

	p[0] = 1
	for i := n; i >= 1; i-- {
		mul1_p()
		if a[i] != 1 {
			convm_p(a[i], rem[i])
		}
	}
	mul1_p()

	fmt.Println(calc_p(rem[0]))
}
