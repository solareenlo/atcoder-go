package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	l := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i], &l[i])
	}

	var mod int
	fmt.Scan(&mod)
	ans := 0
	for i := 0; i < n; i++ {
		t, s := 1, 1
		for ; t <= a[i]; t *= 10 {
		}
		t %= mod
		for ; l[i] > 0; l[i] >>= 1 {
			if l[i]&1 != 0 {
				ans = (ans*t + a[i]*s) % mod
			}
			s *= t + 1
			s %= mod
			t *= t
			t %= mod
		}
	}
	fmt.Println(ans)
}
