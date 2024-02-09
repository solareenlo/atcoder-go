package main

import "fmt"

func main() {
	var c, h int
	fmt.Scan(&c, &h)
	if h >= 2800 {
		fmt.Println("o")
	} else {
		fmt.Println("x")
	}
}
