package main

import "fmt"

func main() {
	var x1, y1, x2, y2 int
	fmt.Scan(&x1, &y1, &x2, &y2)

	dx, dy := x2-x1, y2-y1

	x3, y3 := x2-dy, y2+dx
	x4, y4 := x3-dx, y3-dy

	fmt.Println(x3, y3)
	fmt.Println(x4, y4)
}
