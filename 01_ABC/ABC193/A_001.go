package main

import "fmt"

func main() {
	var a, b float64
	fmt.Scan(&a, &b)

	fmt.Println(100.0 - b/a*100.0)
}
