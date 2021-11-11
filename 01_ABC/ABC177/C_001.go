package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}

	mod := int(1e9 + 7)
	s := make([]int, n+1)
	for i := 0; i < n; i++ {
		s[i+1] += s[i] + a[i]
		s[i+1] %= mod
	}

	res := 0
	for i := 0; i < n-1; i++ {
		tmp := s[n] - s[i+1]
		if tmp < 0 {
			tmp += mod
		}
		res += a[i] * tmp
		res %= mod
	}

	fmt.Println(res)
}
