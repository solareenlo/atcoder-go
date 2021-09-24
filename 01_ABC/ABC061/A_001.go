package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	if c >= a && c <= b {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
