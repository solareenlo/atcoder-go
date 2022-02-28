package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	res := 1 << 60
	for i := 0; i < 61; i++ {
		a := n / (1 << i)
		b := i
		c := n - a*(1<<i)
		res = min(res, a+b+c)
	}
	fmt.Println(res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
