package main

import (
	"fmt"
	"sort"
)

func main() {
	var P int
	fmt.Scan(&P)
	P--

	tot := 0
	a := make([]int, 7007)
	for i := 1; i*i < P+1; i++ {
		if P%i == 0 {
			tot++
			a[tot] = i
			if i*i != P {
				tot++
				a[tot] = P / i
			}
		}
	}
	sort.Ints(a[1 : tot+1])

	mod := 998244353
	res := 1
	ph := make([]int, 7007)
	for i := 1; i < tot+1; i++ {
		ph[i] = a[i]
		for j := 1; j < i; j++ {
			if a[i]%a[j] == 0 {
				ph[i] -= ph[j]
			}
		}
		res = (res + ph[i]%mod*(a[i]%mod)%mod) % mod
	}

	fmt.Println(res)
}
