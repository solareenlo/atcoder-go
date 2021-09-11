package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	res, r := 0, 1
	for l := 0; l < n; l++ {
		for r < n && (r <= l || a[r-1] < a[r]) {
			r++
		}
		res += r - l
	}
	fmt.Println(res)
}
