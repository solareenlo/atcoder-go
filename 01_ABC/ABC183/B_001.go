package main

import "fmt"

func main() {
	var sx, sy, gx, gy float64
	fmt.Scan(&sx, &sy, &gx, &gy)

	fmt.Println(sx + (sy*(gx-sx))/(sy+gy))
}
