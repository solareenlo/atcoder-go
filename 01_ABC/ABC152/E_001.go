package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	res := 0
	cnt := [1000001]int{}
	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		res += invMod(a)
		res %= mod
		for j := 2; j*j < a+1; j++ {
			c := 0
			for a%j == 0 {
				a /= j
				c++
			}
			cnt[j] = max(cnt[j], c)
		}
		cnt[a] = max(cnt[a], 1)
	}

	for i := 1; i < 1000001; i++ {
		for j := 1; j < cnt[i]+1; j++ {
			res = res * i % mod
		}
	}

	fmt.Println(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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

func invMod(a int) int {
	return powMod(a, mod-2)
}
