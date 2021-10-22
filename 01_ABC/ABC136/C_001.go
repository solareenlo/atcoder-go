package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	h := make([]int, n)
	for i := range h {
		fmt.Scan(&h[i])
	}

	maxi := 0
	for i := 0; i < n; i++ {
		if maxi < h[i]-1 {
			maxi = h[i] - 1
		}
		if maxi > h[i] {
			fmt.Println("No")
			return
		}
	}

	fmt.Println("Yes")
}
