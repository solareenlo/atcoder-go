package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n+1)
	s := 0
	for i := 1; i <= n; i++ {
		fmt.Scan(&a[i])
		if a[i] == a[i-1] {
			a[i] = -1
			s++
		}
	}
	fmt.Println(s)
}
