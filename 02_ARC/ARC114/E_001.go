package main

import "fmt"

func main() {
	var H, W, a, b, c, d int
	fmt.Scan(&H, &W, &a, &b, &c, &d)

	A := min(a, c) - 1
	B := H - max(a, c)
	C := min(b, d) - 1
	D := W - max(b, d)
	K := abs(a-c) + abs(b-d)

	ans := 0
	for i := 1; i <= A; i++ {
		ans += divMod(1, i+K)
		ans %= mod
	}
	for i := 1; i <= B; i++ {
		ans += divMod(1, i+K)
		ans %= mod
	}
	for i := 1; i <= C; i++ {
		ans += divMod(1, i+K)
		ans %= mod
	}
	for i := 1; i <= D; i++ {
		ans += divMod(1, i+K)
		ans %= mod
	}
	fmt.Println(ans + 1)
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
