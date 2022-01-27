package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	x := 0
	y := 1
	for i := 0; i < n; i++ {
		z := x + y
		x = y
		y = z
	}

	fmt.Println(x, y)
}
