package main

import "fmt"

func main() {
	var x float64
	fmt.Scan(&x)

	fmt.Println(int(x + 0.5))
}
