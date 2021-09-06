package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	if x < y {
		fmt.Println("Better")
	} else {
		fmt.Println("Worse")
	}
}
