package main

import "fmt"

func main() {
	var x, p, a, b int
	fmt.Scan(&x, &p, &a, &b)
	if a/p < b/p {
		fmt.Println(1)
		return
	}

	x %= p
	a %= p - 1
	b %= p - 1
	if a > b {
		b += p - 1
	}

	n := powMod(x, a-1, p)
	ans := p - 1
	for i := 0; i < b-a+1; i++ {
		n = n * x % p
		ans = min(ans, n)
	}

	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func powMod(a, n, mod int) int {
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
