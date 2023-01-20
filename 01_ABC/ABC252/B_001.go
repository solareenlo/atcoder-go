package main

import (
	"fmt"
	"os"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	maxs := make([]int, 0)

	var max int
	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		if a > max {
			max = a
			maxs = make([]int, 0)
		}
		if a == max {
			maxs = append(maxs, i)
		}
	}

	for i := 0; i < k; i++ {
		var b int
		fmt.Scan(&b)
		for _, v := range maxs {
			if v+1 == b {
				fmt.Println("Yes")
				os.Exit(0)
			}
		}
	}

	fmt.Println("No")
}
