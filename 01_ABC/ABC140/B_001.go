package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n+1)
	b := make([]int, n)
	c := make([]int, n-1)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
		a[i]--
	}
	for i := range b {
		fmt.Scan(&b[i])
	}
	for i := range c {
		fmt.Scan(&c[i])
	}

	sum := 0
	for i := 0; i < n; i++ {
		if a[i] == n-1 {
			sum += b[a[i]]
		} else if a[i]+1 == a[i+1] {
			sum += b[a[i]] + c[a[i]]
		} else {
			sum += b[a[i]]
		}
	}

	fmt.Println(sum)
}
