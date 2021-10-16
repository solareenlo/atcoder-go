package main

import "fmt"

func main() {
	var n, maxi, h int
	fmt.Scan(&n, &maxi)

	cnt := 1
	for i := 0; i < n-1; i++ {
		fmt.Scan(&h)
		if maxi <= h {
			cnt++
		}
		maxi = max(maxi, h)
	}
	fmt.Println(cnt)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
