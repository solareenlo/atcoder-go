package main

import (
	"bufio"
	"fmt"
	"os"
)

const kN = 200005
const kP = 998244353

var fac, inv []int

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	a := make([]int, kN)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}

	fac = make([]int, kN)
	inv = make([]int, kN)
	fac[0] = 1
	inv[0] = 1
	for i := 1; i < kN; i++ {
		fac[i] = fac[i-1] * i % kP
		inv[i] = P(fac[i], kP-2)
	}

	f := make([][]map[int]int, kN)
	for i := range f {
		f[i] = make([]map[int]int, 3)
		for j := range f[i] {
			f[i][j] = make(map[int]int)
		}
	}
	f[1][2][a[1]-1] = 1
	for i := 2; i <= n; i++ {
		for j := 0; j < 3; j++ {
			for r1, r2 := range f[i-1][j] {
				for q := 0; q <= j && q+r1 <= a[i]; q++ {
					r := a[i] - r1 - q
					f[i][q][r] += r2 * C(a[i]-1, r1+q-1) % kP * C(j, q) % kP
					f[i][q][r] %= kP
				}
			}
		}
	}

	fmt.Println((f[n][1][0] + f[n][2][0] + f[n][0][0]) % kP)
}

func P(a, b int) int {
	r := 1
	for b != 0 {
		if (b & 1) != 0 {
			r = r * a % kP
		}
		b = b / 2
		a = a * a % kP
	}
	return r
}

func C(a, b int) int {
	if a < b || a < 0 || b < 0 {
		return 0
	}
	return fac[a] * inv[b] % kP * inv[a-b] % kP
}
