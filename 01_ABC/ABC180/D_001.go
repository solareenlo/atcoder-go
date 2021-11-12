package main

import "fmt"

func main() {
	var x, y, a, b int
	fmt.Scan(&x, &y, &a, &b)

	res := 0
	for y/x > a && x*a <= x+b {
		x *= a
		res++
	}

	res += (y - 1 - x) / b
	fmt.Println(res)
}
