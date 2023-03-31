package main

import (
	"fmt"
)

var Mod int64 = 1000000007

func Pow(a, b int64) int64 {
	r := int64(1)
	a %= Mod
	for b != 0 {
		if (b & 1) != 0 {
			r = r * a % Mod
		}
		a = a * a % Mod
		b >>= 1
	}
	return r
}

func Calc(n, m, K int64) int64 {
	if m < 0 {
		return 0
	}
	if n == 0 {
		if K == 0 {
			return 1
		}
		return 0
	}
	if m == 0 {
		if K == 0 {
			return 1
		}
		return 0
	}
	x := int64(1)
	for x <= m {
		x *= 2
	}
	if x <= K {
		return 0
	}
	r := (Pow(m+1, n) - Pow(Mod-1, n)*Pow(x-m-1, n)%Mod + Mod) * Pow(x, Mod-2) % Mod
	var t int64
	if n%2 == 0 {
		t = Calc(n, x-m-2, K)
	} else {
		t = Calc(n, x-m-2, K^(x-1))
	}
	return (r + Pow(Mod-1, n)*t) % Mod
}

func main() {
	var n, m, K int64
	fmt.Scan(&n, &m, &K)
	fmt.Println(Calc(n, m, K))
}
