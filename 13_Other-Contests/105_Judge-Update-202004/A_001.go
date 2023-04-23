package main

import "fmt"

func main() {
	var s, l, r int
	fmt.Scan(&s, &l, &r)
	fmt.Println(max(l, min(s, r)))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
