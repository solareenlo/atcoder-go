package main

import "fmt"

var mod int = 1e9 + 7
var p [1001][1001]int = [1001][1001]int{}
var r, c, x, y, d, l int

func f(a, b int) int {
	if a >= 0 && b >= 0 {
		return p[a*b][d+l]
	}
	return 0
}

func main() {
	p[0][0] = 1
	for i := 1; i < 1000; i++ {
		p[i][0] = 1
		p[i][i] = 1
		for j := 1; j < i; j++ {
			p[i][j] = (p[i-1][j-1] + p[i-1][j]) % mod
		}
	}
	fmt.Scan(&r, &c, &x, &y, &d, &l)
	res := (f(x, y) - (2*f(x-1, y) + 2*f(x, y-1)) + (4*f(x-1, y-1) + f(x-2, y) + f(x, y-2)) - (2*f(x-2, y-1) + 2*f(x-1, y-2)) + f(x-2, y-2)) % mod
	res = (res + mod) % mod
	res = res * (r - x + 1) % mod * (c - y + 1) % mod * p[d+l][d] % mod
	fmt.Println(res)
}
