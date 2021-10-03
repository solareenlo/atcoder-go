package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	if a+b > c+d {
		fmt.Println("Left")
	} else if a+b < c+d {
		fmt.Println("Right")
	} else {
		fmt.Println("Balanced")
	}
}
