package main

import "fmt"

func main() {
	var n, x, y int
	fmt.Scan(&n, &x, &y)

	b, r := 0, 1
	for i := n; i > 1; i-- {
		b += r * x
		r += b
		b *= y
	}

	fmt.Println(b)
}
