package main

import "fmt"

func main() {
	var n, x int
	fmt.Scan(&n, &x)
	a := make([]int, n)
	s := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
		if a[i] <= x {
			s += a[i]
		}
	}
	fmt.Println(s)
}
