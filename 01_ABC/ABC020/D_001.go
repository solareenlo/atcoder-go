package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	d := make([]int, 0)
	for i := 1; i*i <= k; i++ {
		if k%i == 0 {
			d = append(d, i)
			if i*i != k {
				d = append(d, k/i)
			}
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(d)))

	mod, r := int(1e9+7), 0
	s := make([]int, len(d))
	for i := 0; i < len(d); i++ {
		x := d[i]
		s[i] = (x + (n - n%x)) * (n / x) / 2 % mod
		for j := 0; j < i; j++ {
			if d[j]%d[i] == 0 {
				s[i] -= s[j]
			}
		}
		r += (s[i]%mod + mod) % mod * k / x
	}
	fmt.Println(r % mod)
}
