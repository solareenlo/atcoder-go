package main

import "fmt"

func main() {
	var x, y, z int
	fmt.Scan(&x, &y, &z)

	fmt.Println((y*z - 1) / x)
}
