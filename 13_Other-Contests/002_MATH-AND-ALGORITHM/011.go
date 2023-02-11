package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	m := Eratos(n)

	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for i := range keys {
		fmt.Print(keys[i], " ")
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
