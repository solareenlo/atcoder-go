package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	diff := 2025 - n
	for i := 1; i < 10; i++ {
		if diff%i == 0 && diff/i < 10 {
			fmt.Println(i, "x", diff/i)
		}
	}
}
