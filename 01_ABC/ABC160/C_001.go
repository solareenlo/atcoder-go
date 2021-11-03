package main

import "fmt"

func main() {
	var k, n int
	fmt.Scan(&k, &n)

	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}

	maxi := 0
	for i := 0; i < n-1; i++ {
		maxi = max(maxi, abs(a[i]-a[i+1]))
	}

	maxi = max(maxi, k-a[n-1]+a[0])

	fmt.Println(k - maxi)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
