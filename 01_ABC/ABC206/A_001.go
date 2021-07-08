package main

import "fmt"

func main() {
	var n float64
	fmt.Scan(&n)
	m := int(n * 1.08)
	if m == 206 {
		fmt.Println("so-so")
	} else if m < 206 {
		fmt.Println("Yay!")
	} else {
		fmt.Println(":(")
	}
}
