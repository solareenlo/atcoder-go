package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&MOD, &N)

	var dt [10000]int
	for i := 0; i < MOD; i++ {
		dt[powMod(i, N)]++
	}
	ans := 0
	for i := 0; i < MOD; i++ {
		ans += dt[i] * (dt[(i+1)%MOD]*(MOD-1) + dt[i])
	}
	fmt.Println(ans)
}

var MOD int

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
