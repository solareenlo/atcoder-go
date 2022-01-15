package main

import "fmt"

func main() {
	a := [8]int{}
	for i := 0; i < 7; i++ {
		fmt.Scan(&a[i])
	}

	res := 0
	for i := 0; i < 7; i++ {
		var x int
		fmt.Scan(&x)
		res += max(a[i], x)
	}

	fmt.Println(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
