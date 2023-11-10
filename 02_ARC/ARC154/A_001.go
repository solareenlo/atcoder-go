package main

import (
	"fmt"
	"strconv"
)

func main() {
	const MOD = 998244353

	var n int
	var A, B string
	fmt.Scan(&n, &A, &B)
	am, bm := 0, 0
	for i := 0; i < n; i++ {
		a, _ := strconv.Atoi(string(A[i]))
		b, _ := strconv.Atoi(string(B[i]))
		if a > b {
			a, b = b, a
		}
		am = (am*10 + a) % MOD
		bm = (bm*10 + b) % MOD
	}
	fmt.Println((am * bm) % MOD)
}
