package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	c := 1
	for c*c < N {
		c++
	}
	for N%c != 0 {
		c--
	}
	fmt.Println((c + N/c) * 2)
}
