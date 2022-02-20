package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	x := (a + b) / 2
	y := a - x
	fmt.Println(x, y)
}
