package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	if b < c || d < a {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}
