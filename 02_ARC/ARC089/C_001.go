package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	t, x, y := make([]int, n), make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&t[i], &x[i], &y[i])
		if t[i]%2 != (x[i]+y[i])%2 || t[i] < x[i]+y[i] {
			fmt.Println("No")
			return
		}
	}
	for i := 0; i < n-1; i++ {
		if t[i+1]-t[i] < abs(x[i+1]-x[i])+abs(y[i+1]-y[i]) {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
