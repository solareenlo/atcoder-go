package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, 1<<n)
	for i := range a {
		fmt.Scan(&a[i])
	}

	maxi, index := 0, 0
	for i := 0; i < 1<<n; i++ {
		if maxi < a[i] {
			maxi = a[i]
			index = i
		}
	}

	half := 1 << (n - 1)
	start := 0
	if index < half {
		start = half
	}

	maxi = 0
	for i := start; i < start+half; i++ {
		if maxi < a[i] {
			maxi = a[i]
			index = i
		}
	}

	fmt.Println(index + 1)
}
