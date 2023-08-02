package main

import "fmt"

func main() {
	const pai = 3.1415926535
	var r, d float64
	fmt.Scan(&r, &d)
	fmt.Println(r * r * d * 2 * pai * pai)
}
