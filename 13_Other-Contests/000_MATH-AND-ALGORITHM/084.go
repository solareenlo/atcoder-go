package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if c-a-b < 0 {
		fmt.Println("No")
	} else if 4*a*b < (c-a-b)*(c-a-b) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
