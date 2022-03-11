package main

import "fmt"

const mod = 998244353

func add(x, y int) int {
	if x+y >= mod {
		return x + y - mod
	}
	return x + y
}

func main() {
	var n, x int
	fmt.Scan(&n, &x)

	C := [3005][3005]int{}
	C[0][0] = 1
	for i := 1; i <= n; i++ {
		C[i][0] = 1
		for j := 1; j <= i; j++ {
			C[i][j] = add(C[i-1][j], C[i-1][j-1])
		}
	}

	ans := 0
	for i := 0; i < x; i += 2 {
		for j := min(n-i, x-1-i); j > 0 && x-1-i < 2*j; j-- {
			ans = (ans + C[j][x-1-i-j]*C[n][i+j]) % mod
		}
	}

	for i := 0; i <= n; i++ {
		for j := 0; j <= i && i+j < x-1; j++ {
			ans = (ans + C[i][j]*C[n][i]) % mod
		}
	}

	if x&1 != 0 {
		for i := (x - 1) / 2; i <= n; i++ {
			ans = add(ans, C[n][i])
		}
	}
	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
