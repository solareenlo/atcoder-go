package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	x := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&x[i])
	}

	mini := 1 << 62
	for i := 0; i < n-k+1; i++ {
		l := abs(x[i]) + abs(x[k-1+i]-x[i])
		r := abs(x[k-1+i]) + abs(x[k-1+i]-x[i])
		mini = min(mini, min(l, r))
	}
	fmt.Println(mini)
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
