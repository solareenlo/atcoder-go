package main

import "fmt"

var (
	x1, x2, x3, x4, x5, x6 int
	y1, y2, y3, y4, y5, y6 int
)

func f(x, y int) int {
	x++
	y++
	return fact[x+y] * invf[x] % mod * invf[y] % mod
}

func cal(xa, ya, xb, yb int) int {
	return (f(xa-x1, ya-y1) - f(xa-x1, ya-y2) - f(xa-x2, ya-y1) + f(xa-x2, ya-y2)) *
		(f(x6-xb, y6-yb) - f(x5-xb, y6-yb) - f(x6-xb, y5-yb) + f(x5-xb, y5-yb)) % mod
}

func main() {
	initMod()

	fmt.Scan(&x1, &x2, &x3, &x4, &x5, &x6, &y1, &y2, &y3, &y4, &y5, &y6)
	x2++
	y2++
	x5--
	y5--

	ans := 0
	for i := x3; i <= x4; i++ {
		ans -= cal(i, y3-1, i, y3) * (y3 + i)
		ans += cal(i, y4, i, y4+1) * (y4 + i + 1)
	}

	for i := y3; i <= y4; i++ {
		ans -= cal(x3-1, i, x3, i) * (x3 + i)
		ans += cal(x4, i, x4+1, i) * (x4 + i + 1)
	}

	fmt.Println((ans%mod + mod) % mod)
}

const mod = 1000000007
const size = 2000200

var fact, invf [size]int

func initMod() {
	fact[0] = 1
	invf[0] = 1
	for i := int(1); i < size; i++ {
		fact[i] = (fact[i-1] * i) % mod
		invf[i] = invMod(fact[i])
	}
}

func powMod(a, n int) int {
	res := int(1)
	for n > 0 {
		if n%2 == 1 {
			res = res * a % mod
		}
		a = a * a % mod
		n /= 2
	}
	return res
}

func invMod(a int) int {
	return powMod(a, mod-2)
}
