package main

import "fmt"

func main() {
	var r1, r2 float64
	fmt.Scan(&r1, &r2)
	fmt.Println(1 / (1/r1 + 1/r2))
}
