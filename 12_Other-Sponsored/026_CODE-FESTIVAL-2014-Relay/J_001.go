package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	if b >= a/2 || a%2 == 1 {
		fmt.Println("first")
	} else {
		fmt.Println("second")
	}
}
