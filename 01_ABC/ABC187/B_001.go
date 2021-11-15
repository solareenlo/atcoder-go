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

	cnt := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if abs(y[i]-y[j]) <= abs(x[i]-x[j]) {
				cnt++
			}
		}
	}

	fmt.Println(cnt)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
