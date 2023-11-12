package main

import "fmt"

func main() {
	var N, M int
	fmt.Scan(&N, &M)

	n := N
	for M > 0 {
		if n > int(1e9) {
			break
		}
		fmt.Print("o")
		n *= N
		M--
	}
	for M > 0 {
		fmt.Print("x")
		M--
	}
}
