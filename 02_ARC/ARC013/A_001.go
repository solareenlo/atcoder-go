package main

import (
	"fmt"
)

func main() {
	var n, m, l int
	fmt.Scan(&n, &m, &l)

	b := make([]int, 3)
	fmt.Scan(&b[0], &b[1], &b[2])

	maxi := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				if i != j && i != k && j != k {
					maxi = max(maxi, (n/b[i])*(m/b[j])*(l/b[k]))
				}
			}
		}
	}
	fmt.Println(maxi)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
