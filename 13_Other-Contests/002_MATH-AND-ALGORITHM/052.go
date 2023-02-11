package main

import "fmt"

func main() {
	var X, Y int
	fmt.Scan(&X, &Y)

	A := Y - (X+Y)/3
	B := Y - 2*A
	if (X+Y)%3 > 0 || A < 0 || B < 0 {
		fmt.Println(0)
		return
	}

	ans := 1
	for i := 0; i < A; i++ {
		ans *= (A + B - i)
		ans %= mod
	}

	for i := 1; i <= A; i++ {
		ans = divMod(ans, i)
	}
	fmt.Println(ans)
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
