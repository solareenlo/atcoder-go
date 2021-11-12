package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	x := make([]int, n)
	y := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&x[i], &y[i])
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				dx := x[j] - x[i]
				dy := y[j] - y[i]
				if dx*(y[k]-y[i]) == dy*(x[k]-x[i]) {
					fmt.Println("Yes")
					return
				}
			}
		}
	}
	fmt.Println("No")
}
