package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}

	sum := 0
	for i := 0; i < n-1; i++ {
		if a[i+1] < a[i] {
			sum += a[i] - a[i+1]
			a[i+1] = a[i]
		}
	}

	fmt.Println(sum)
}
