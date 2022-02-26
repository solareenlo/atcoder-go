package main

import "fmt"

const P = 998244353
const M = 170

var f = [M]int{}

func qur(x int) int {
	res := 1
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			cnt := 0
			for x%i == 0 {
				cnt++
				x /= i
			}
			res = res * f[cnt] % P
		}
	}
	if x > 1 {
		res = res * f[1] % P
	}
	return res
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	f[0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j < M; j++ {
			f[j] = (f[j] + f[j-1]) % P
		}
	}

	ans := 0
	for i := 1; i <= m; i++ {
		ans = (ans + qur(i)) % P
	}
	fmt.Println(ans)
}
