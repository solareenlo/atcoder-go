package main

import "fmt"

func main() {
	var n, m, x, y int
	fmt.Scan(&n, &m, &x, &y)

	ans := 4
	if x == 1 {
		ans--
	}
	if x == n {
		ans--
	}
	if y == 1 {
		ans--
	}
	if y == m {
		ans--
	}
	fmt.Println(ans)
}
