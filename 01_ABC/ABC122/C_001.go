package main

import "fmt"

func main() {
	var n, q int
	var s string
	fmt.Scan(&n, &q, &s)

	ac := make([]int, n)
	cnt := 0
	for i := 0; i < n-1; i++ {
		ac[i] = cnt
		if s[i] == 'A' && s[i+1] == 'C' {
			cnt++
		}
	}
	ac[n-1] = cnt

	var l, r int
	for i := 0; i < q; i++ {
		fmt.Scan(&l, &r)
		fmt.Println(ac[r-1] - ac[l-1])
	}
}
