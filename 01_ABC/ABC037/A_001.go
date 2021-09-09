package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	if a < b {
		fmt.Println(c / a)
	} else {
		fmt.Println(c / b)
	}
}
