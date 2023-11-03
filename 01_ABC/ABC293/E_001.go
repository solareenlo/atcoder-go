package main

import "fmt"

func solve(a, x, m int) int {
	if x == 0 {
		return 0
	}
	if x%2 == 1 {
		return (solve(a, x-1, m)*a + 1) % m
	}
	if x%2 == 0 {
		return (solve(a*a%m, x/2, m) * (1 + a)) % m
	}
	return 0
}

func main() {
	var a, x, m int
	fmt.Scan(&a, &x, &m)
	fmt.Println(solve(a, x, m))
}
