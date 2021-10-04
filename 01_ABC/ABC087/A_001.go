package main

import "fmt"

func main() {
	var x, a, b int
	fmt.Scan(&x, &a, &b)

	fmt.Println((x - a) % b)
}
