package main

import "fmt"

func main() {
	var n, m, t int
	fmt.Scan(&n, &m, &t)

	s, battery := 0, n
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		battery -= a - s
		if battery <= 0 {
			fmt.Println("No")
			return
		}
		battery = min(n, battery+b-a)
		s = b
	}

	battery -= t - s
	if battery <= 0 {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
