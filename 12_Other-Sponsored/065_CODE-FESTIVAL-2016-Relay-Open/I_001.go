package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n == 2 {
		fmt.Println(-1)
		return
	}
	for i := 0; i < n; i++ {
		a := make([]int, n)
		for j := 0; j < n; j++ {
			a[j] = (i+j+1)%n + 1
		}
		if (^n&1) != 0 && i >= n>>1 {
			a[(n>>1)-1], a[n>>1] = a[n>>1], a[(n>>1)-1]
		}
		for j := 0; j < n-1; j++ {
			fmt.Print(a[j], " ")
		}
		fmt.Println()
	}
}
