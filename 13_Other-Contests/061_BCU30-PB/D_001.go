package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, q int
	fmt.Scan(&n, &q)

	var x = make([]int, n)
	var s = make([]int, n+1)
	for i := 0; i < n; i++ {
		fmt.Scan(&x[i])
		s[i+1] = s[i] + x[i]
	}
	sort.Ints(x)

	for i := 0; i < q; i++ {
		var t int
		fmt.Scan(&t)
		m := sort.Search(len(x), func(i int) bool { return t < x[i] })
		fmt.Println(s[n]-s[m]-t*(n-m)+t*m-s[m])
	}
}
