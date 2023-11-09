package main

import "fmt"

func main() {
	var n, m, p int
	fmt.Scan(&n, &m, &p)
	if m > n {
		fmt.Println(0)
	} else {
		fmt.Println(max(0, (n-m)/p+1))
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
