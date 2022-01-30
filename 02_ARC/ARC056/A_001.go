package main

import "fmt"

func main() {
	var a, b, k, l int
	fmt.Scan(&a, &b, &k, &l)

	ans1 := ((k / l) + 1) * b
	ans2 := (k/l)*b + (k%l)*a
	ans3 := a * k

	ans1 = min(ans1, ans2)
	ans1 = min(ans1, ans3)
	fmt.Println(ans1)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
