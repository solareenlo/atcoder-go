package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	for i := 9; i >= 0; i-- {
		pos := (1 << i)
		fmt.Print((N / pos) % 2)
	}
	fmt.Println()
}
