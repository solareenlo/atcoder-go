package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if (a < c && b > c) || (a > c && b < c) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
