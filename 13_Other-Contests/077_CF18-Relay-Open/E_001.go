package main

import (
	"bufio"
	"fmt"
	"os"
)

const MOD = 1000000007

var fact, ifact [1000001]int

func main() {
	in := bufio.NewReader(os.Stdin)

	var S string
	var A int
	fmt.Fscan(in, &S, &A)
	N := len(S)
	x := A
	y := N - A
	for _, c := range S {
		if c == 'W' {
			x--
		} else if c == 'F' {
			y--
		}
	}
	fact[0] = 1
	for i := 0; i < N; i++ {
		fact[i+1] = (i + 1) * fact[i] % MOD
	}
	ifact[N] = invMod(fact[N])
	for i := N - 1; i >= 0; i-- {
		ifact[i] = ifact[i+1] * (i + 1) % MOD
	}
	ans := 0
	for i := 0; i < N; i++ {
		c1 := S[i]
		c2 := S[(i+1)%N]
		if c1 > c2 {
			c1, c2 = c2, c1
		}
		if c1 == '?' {
			if c2 == '?' {
				if x >= 1 && y >= 1 {
					ans += C(x+y-2, x-1) * 2 % MOD
				}
			} else if c2 == 'F' {
				if x >= 1 {
					ans += C(x+y-1, y)
				}
			} else {
				if y >= 1 {
					ans += C(x+y-1, x)
				}
			}
		} else if c1 == 'F' && c2 == 'W' {
			ans += C(x+y, x)
		}
		ans %= MOD
	}
	fmt.Println(ans)
}

func C(n, m int) int {
	return fact[n] * ifact[m] % MOD * ifact[n-m] % MOD
}

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

func invMod(a int) int {
	return powMod(a, MOD-2)
}
