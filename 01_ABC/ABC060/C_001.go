package main

import "fmt"

func main() {
	var n, T int
	fmt.Scan(&n, &T)
	t := make([]int, n)
	for i := range t {
		fmt.Scan(&t[i])
	}

	res := 0
	for i := 0; i < n-1; i++ {
		res += min(t[i+1]-t[i], T)
	}
	res += T
	fmt.Println(res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
