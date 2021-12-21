package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n >= 42 {
		fmt.Printf("AGC0%02d", n+1)
	} else {
		fmt.Printf("AGC0%02d", n)
	}
	fmt.Println()
}
