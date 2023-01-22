package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var a [100]int
	a[1] = 1
	for i := 1; i <= n; i++ {
		for j := i; j >= 1; j-- {
			a[j] += a[j-1]
			fmt.Printf("%d ", a[j])
		}
		fmt.Println()
	}
}
