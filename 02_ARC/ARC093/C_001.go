package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n+2)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i+1])
	}

	sum := 0
	for i := 0; i < n+1; i++ {
		sum += abs(a[i] - a[i+1])
	}

	for i := 1; i < n+1; i++ {
		l := min(a[i-1], a[i+1])
		r := max(a[i-1], a[i+1])
		if l <= a[i] && a[i] <= r {
			fmt.Println(sum)
		} else if a[i] < l {
			fmt.Println(sum - (l-a[i])*2)
		} else if a[i] > r {
			fmt.Println(sum - (a[i]-r)*2)
		}
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
