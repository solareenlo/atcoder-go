package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	v := [1000002]int{}
	for i := 0; i < n; i++ {
		var l, r int
		fmt.Scan(&l, &r)
		v[l]++
		v[r+1]--
	}

	s := [1000003]int{}
	for i := range v {
		s[i+1] = s[i] + v[i]
	}

	res := 0
	for i := range s {
		res = max(res, s[i])
	}
	fmt.Println(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
