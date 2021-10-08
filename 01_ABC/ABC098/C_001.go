package main

import "fmt"

func main() {
	var n int
	var s string
	fmt.Scan(&n, &s)

	sumW := make([]int, n+1)
	for i := 0; i < n; i++ {
		if s[i] == 'W' {
			sumW[i+1] = sumW[i] + 1
		} else {
			sumW[i+1] = sumW[i]
		}
	}
	sumE := make([]int, n+1)
	for i := 0; i < n; i++ {
		if s[i] == 'E' {
			sumE[i+1] = sumE[i] + 1
		} else {
			sumE[i+1] = sumE[i]
		}
	}

	mini := 1 << 60
	for i := 0; i < n+1; i++ {
		r := sumW[i]
		l := sumE[n] - sumE[i]
		mini = min(mini, r+l)
	}
	fmt.Println(mini)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
