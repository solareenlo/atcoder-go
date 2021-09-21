package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)
	v := make([]int, n)
	for i := range v {
		fmt.Scan(&v[i])
	}
	sort.Sort(sort.Reverse(sort.IntSlice(v)))

	ave := 0.0
	for i := 0; i < a; i++ {
		ave += float64(v[i])
	}
	ave /= float64(a)

	l, r := a-1, a-1
	for l-1 >= 0 && v[l-1] == v[a-1] {
		l--
	}
	for r+1 < n && v[r+1] == v[a-1] {
		r++
	}

	if l != 0 {
		n = a
	} else {
		n = b
	}
	res := 0
	for i := a; i <= n; i++ {
		res += nCr(r-l+1, i-l)
	}
	fmt.Printf("%f\n%d\n", ave, res)
}

func nCr(n, r int) int {
	if n < r || n < 0 || r < 0 {
		return 0
	}
	res := 1
	for i := 1; i <= r; i++ {
		res *= n + 1 - i
		res /= i
	}
	return res
}
