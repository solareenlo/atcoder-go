package main

import "fmt"

const mod = 1_000_000_007

func pow10(x int) int {
	res := 1
	tmp := 10
	for x != 0 {
		if x&1 != 0 {
			res = res * tmp % mod
		}
		tmp = tmp * tmp % mod
		x >>= 1
	}
	return res
}

func main() {
	var n int
	fmt.Scan(&n)

	f := make([]int, 23000001)
	for i := 1; i <= 23000000; i++ {
		f[i] = f[i/10] + 1
	}

	ans := 0
	for i, j, s := 1, 1, 0; i < 1e7; i++ {
		for s < n {
			s += f[j]
			j++
		}
		if s == n {
			ans++
		}
		s -= f[i]
	}

	for i := 1; i <= n/8; i++ {
		if n%i != 0 {
			ans++
		} else {
			ans = (ans + 9*pow10(n/i-1) - i + 1) % mod
		}
	}
	fmt.Println(ans)
}
