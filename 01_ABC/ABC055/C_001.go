package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	res := m / 2
	if 2*n < m {
		res = n + ((m - 2*n) / 4)
	}
	fmt.Println(res)
}
