package main

import "fmt"

func main() {
	size := int(1e5)
	primes := Eratos(size + 1)
	p := make([]int, size+1)
	for i := 3; i < size; i += 2 {
		j := (i + 1) / 2
		if primes[i] && primes[j] {
			p[i] = 1
		}
	}

	s := make([]int, size+2)
	for i := 1; i < size+1; i++ {
		s[i] = s[i-1] + p[i]
	}

	var q, l, r int
	fmt.Scan(&q)
	for i := 0; i < q; i++ {
		fmt.Scan(&l, &r)
		fmt.Println(s[r] - s[l-1])
	}
}

func Eratos(n int) map[int]bool {
	primes := make(map[int]bool)
	f := make([]int, n+1)
	f[0], f[1] = -1, -1
	for i := 2; i <= n; i++ {
		if f[i] != 0 {
			continue
		}
		primes[i] = true
		f[i] = i
		for j := i * i; j <= n; j += i {
			if f[j] == 0 {
				f[j] = i
			}
		}
	}
	return primes
}
