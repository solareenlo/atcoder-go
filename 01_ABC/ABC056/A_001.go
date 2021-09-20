package main

import "fmt"

func main() {
	var a, b string
	fmt.Scan(&a, &b)

	var i, j int
	if a == "H" {
		i = 1
	} else {
		i = -1
	}
	if b == "H" {
		j = 1
	} else {
		j = -1
	}
	if i*j == 1 {
		fmt.Println("H")
	} else {
		fmt.Println("D")
	}
}
