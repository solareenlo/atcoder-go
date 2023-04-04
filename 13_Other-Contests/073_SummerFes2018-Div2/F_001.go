package main

import "fmt"

func main() {
	var a, mod, k int
	fmt.Scan(&a, &mod, &k)

	ans := 0
	k--
	for i, p := 0, -1; i < k; i++ {
		x := mtetra(a, i, mod) % mod
		if p == x {
			ans += (k - i) % mod * x % mod
			ans %= mod
			break
		}
		ans += x
		ans %= mod
		p = x
	}
	fmt.Println(ans)
}

func mtetra(a, n, m int) int {
	if m == 1 {
		return 0
	}
	if n == 0 {
		return 1
	}
	if n == 1 {
		return a
	}
	return mpow(a, mtetra(a, n-1, phi(m)), m)
}

func mpow(x, n, m int) int {
	if x == 1 || n == 0 {
		return 1
	}
	v := 1
	flg := x >= m
	x %= m
	for n > 0 {
		if (n & 1) != 0 {
			v *= x
		}
		if v >= m {
			flg = true
			v %= m
		}
		x = x * x % m
		n >>= 1
	}
	if flg {
		return v + m
	}
	return v
}

func phi(n int) int {
	res := n
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			res = res / i * (i - 1)
			for n%i == 0 {
				n /= i
			}
		}
	}
	if n != 1 {
		res = res / n * (n - 1)
	}
	return res
}
