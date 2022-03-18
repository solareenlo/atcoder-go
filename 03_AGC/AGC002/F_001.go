package main

import "fmt"

const mod = 1_000_000_007

var (
	a = make([]int, 4000004)
	b = make([]int, 4000004)
)

func c(n, m int) int {
	return b[m] * b[n-m] % mod * a[n] % mod
}

func main() {
	var m, n int
	fmt.Scan(&m, &n)
	b[1] = 1
	f := make([]int, 2002)
	f[0] = 1
	n--
	if n == 0 {
		fmt.Println(1)
		return
	}

	for i := 2; i <= m*(n+1); i++ {
		b[i] = (mod - mod/i) * b[mod%i] % mod
	}

	a[0] = 1
	b[0] = 1
	for i := 1; i <= m*(n+1); i++ {
		a[i] = a[i-1] * i % mod
		b[i] = b[i-1] * b[i] % mod
	}

	for i := 0; i < m; i++ {
		for j := i; j < m; j++ {
			f[j+1] = (f[j+1] + f[j]*c(i+(j+1)*n-1, n-1)) % mod
		}
	}

	fmt.Println(f[m] * a[m] % mod)
}
