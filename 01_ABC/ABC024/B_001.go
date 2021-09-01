package main

import "fmt"

func main() {
	var n, t int
	fmt.Scan(&n, &t)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	res := 0
	for i := 0; i < n; i++ {
		if i == n-1 {
			continue
		}
		if a[i+1]-a[i] >= t {
			res += t
		} else {
			res += a[i+1] - a[i]
		}
	}
	res += t
	fmt.Println(res)
}
