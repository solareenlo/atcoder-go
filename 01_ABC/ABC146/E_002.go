package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}

	s := make([]int, n+1)
	s[0] = 0
	for i := 0; i < n; i++ {
		s[i+1] = (s[i] + a[i] - 1) % k
	}

	res := 0
	c := map[int]int{}
	for i, e := range s {
		res += c[e]
		c[e]++
		if i-(k-1) >= 0 {
			c[s[i-k+1]]--
		}
	}

	fmt.Println(res)
}
