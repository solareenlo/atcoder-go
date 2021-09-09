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
	for i := 0; i < n; i++ {
		s[i+1] = s[i] + a[i]
	}

	res := 0
	for i := 0; i < n; i++ {
		if i+n-k+1 <= n {
			res += s[i+n-k+1] - s[i]
		}
	}
	fmt.Println(res)
}
