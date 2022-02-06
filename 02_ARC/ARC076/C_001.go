package main

import "fmt"

const mod int = int(1e9 + 7)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	res, diff := 1, abs(n-m)
	if diff > 1 {
		res = 0
	} else {
		if diff == 0 {
			res = cal(res, n)
			res = cal(res, m)
			res = (res * 2) % mod
		}
		if diff == 1 {
			res = cal(res, n)
			res = cal(res, m)
		}
	}
	fmt.Println(res)
}

func cal(res, n int) int {
	for i := 0; i < n; i++ {
		res *= i + 1
		res %= mod
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
