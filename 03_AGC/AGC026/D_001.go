package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	h := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&h[i])
	}

	h[0] = 1
	f := [101][101]int{}
	f[0][0] = 1
	for i := 1; i <= n; i++ {
		for j := 0; j <= n; j++ {
			if h[j] > h[i] {
				continue
			}
			if h[j] < h[i-1] {
				if h[i] < h[i-1] {
					f[i][j] = (f[i-1][j] + f[i-1][i]) % mod
				} else {
					f[i][j] = (f[i-1][j] + f[i-1][i-1]) * powMod(2, h[i]-h[i-1]) % mod
				}
			} else {
				f[i][j] = 2 * f[i-1][i-1] * powMod(2, h[i]-h[j]) % mod
			}
		}
	}

	fmt.Println(f[n][0])
}

const mod = 1000000007

func powMod(a, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = res * a % mod
		}
		a = a * a % mod
		n /= 2
	}
	return res
}
