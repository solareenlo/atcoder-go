package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	initMod()

	var N, A, B, C int
	fmt.Fscan(in, &N, &A, &B, &C)

	if ((B + C + N) & 1) != 0 {
		fmt.Println(0)
		return
	}

	Y := (N - C + B) / 2
	X := N - Y

	ways := make([]int, A+2)
	for pos := 0; pos <= A+1; pos++ {
		offset := B + A - pos + 1
		ways[pos] = calc(-offset, offset, X, Y)
	}

	ans := 0
	for pos := A; pos >= 0; pos-- {
		cur := (ways[pos+1] - ways[pos] + MOD) % MOD
		if cur == 0 {
			continue
		}
		cur = (cur * p2[X-(C+A-pos)]) % MOD
		ans = (ans + cur) % MOD
	}
	fmt.Println(ans)
}

func calc(x1, y1, x2, y2 int) int {
	x := x2 - x1
	y := y2 - y1
	if x < 0 || y < 0 {
		return 0
	}
	return nCrMod(x+y, x)
}

const MOD = 998244353
const size = 10000000

var fact, invf, p2 [size + 5]int

func initMod() {
	fact[0] = 1
	invf[0] = 1
	for i := int(1); i <= size; i++ {
		fact[i] = (fact[i-1] * i) % MOD
		invf[i] = invMod(fact[i])
	}
	p2[0] = 1
	for i := 1; i <= size; i++ {
		p2[i] = (p2[i-1] + p2[i-1]) % MOD
	}
}

func powMod(a, n int) int {
	res := int(1)
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

func nCrMod(n, r int) int {
	if n < r || n < 0 || r < 0 {
		return 0
	}
	return fact[n] * invf[r] % MOD * invf[n-r] % MOD
}
