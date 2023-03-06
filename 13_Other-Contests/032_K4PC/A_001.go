package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&a[i])
	}
	sum := 0
	for i := 1; i <= n; i++ {
		sum += abs(a[i] - i)
	}
	fmt.Println(sum / 2)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
