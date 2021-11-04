package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n+1)
	for i := 1; i < n; i++ {
		var boss int
		fmt.Scan(&boss)
		a[boss]++
	}

	for i := 1; i < n+1; i++ {
		fmt.Println(a[i])
	}
}
