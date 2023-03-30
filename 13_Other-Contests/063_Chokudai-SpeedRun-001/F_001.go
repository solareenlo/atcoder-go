package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var a, m, res int
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		if a > m {
			res++
			m = a
		}
	}
	fmt.Println(res)
}
