package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	res := make([]int, n)
	var a, b int
	for i := 0; i < m; i++ {
		fmt.Scan(&a, &b)
		res[a-1]++
		res[b-1]++
	}

	for i := 0; i < n; i++ {
		fmt.Println(res[i])
	}
}
