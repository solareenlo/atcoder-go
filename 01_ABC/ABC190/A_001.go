package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if c != 0 {
		b--
	}
	if a > b {
		fmt.Println("Takahashi")
	} else {
		fmt.Println("Aoki")
	}
}
