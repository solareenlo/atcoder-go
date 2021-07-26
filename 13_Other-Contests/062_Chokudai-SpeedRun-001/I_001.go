package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}

	r, sum, res := 0, 0, 0
	for l := 0; l < n; l++ {
		for r < n && sum < n {
			sum += a[r]
			r++
		}
		if sum == n {
			res++
		}
		sum -= a[l]
	}
	fmt.Println(res)
}
