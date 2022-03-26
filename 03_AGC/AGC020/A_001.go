package main

import "fmt"

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)

	if (b-a)%2 == 1 {
		fmt.Println("Borys")
	} else {
		fmt.Println("Alice")
	}
}
