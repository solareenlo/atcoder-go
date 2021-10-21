package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	sum, sum0 := 0, 0
	for i := range a {
		fmt.Scan(&a[i])
		sum += a[i]
		if i%2 != 0 {
			sum0 += a[i]
		}
	}

	res := make([]int, n)
	res[0] = sum - 2*sum0
	for i := 0; i < n-1; i++ {
		res[i+1] = 2*a[i] - res[i]
	}
	for i := 0; i < n; i++ {
		fmt.Print(res[i])
		if i != n-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
