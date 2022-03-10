package main

import "fmt"

func main() {
	var h, w, k int
	fmt.Scan(&h, &w, &k)

	const mod = 1_000_000_007
	n := h + w
	a := 1
	if k == 1 {
		fmt.Println(2 * n)
	} else {
		a = (k+1)*(k-1)%mod*k%mod*333333336%mod*h%mod*w%mod +
			(k)*(k+3)%mod*n%mod*(n-1)%mod*500000004%mod
		a %= mod
	}

	for i := n - 2; i >= n-k+1; i-- {
		a = a * i % mod
	}
	fmt.Println(a)
}
