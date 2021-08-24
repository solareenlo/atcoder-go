package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	a, b, c := -1, -1, -1
	diff := m - (3 * n)
	if diff == 0 {
		a = 0
		c = 0
		b = n
	} else if 0 < diff && diff <= n {
		a = 0
		b = n - diff
		c = diff
	} else if abs(diff) <= n {
		a = abs(diff)
		b = n - abs(diff)
		c = 0
	}
	fmt.Println(a, b, c)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
