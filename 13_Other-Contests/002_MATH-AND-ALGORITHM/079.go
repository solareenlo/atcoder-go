package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	fmt.Println(n * (n - 1) / 2)
}
