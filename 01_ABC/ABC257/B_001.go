package main

import "fmt"

func main() {
	var n, k, q int
	fmt.Scan(&n, &k, &q)

	a := make([]int, 205)
	for i := 0; i < k; i++ {
		fmt.Scan(&a[i+1])
	}

	var x int
	for i := 0; i < q; i++ {
		fmt.Scan(&x)
		if a[x+1] != a[x]+1 {
			if a[x] < n {
				a[x]++
			}
		}
	}

	for i := 1; i <= k; i++ {
		fmt.Print(a[i], " ")
	}
}
