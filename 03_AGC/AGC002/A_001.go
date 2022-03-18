package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	if a <= 0 && 0 <= b {
		fmt.Println("Zero")
	} else if a > 0 || (b-a)%2 == 1 {
		fmt.Println("Positive")
	} else {
		fmt.Println("Negative")
	}
}
