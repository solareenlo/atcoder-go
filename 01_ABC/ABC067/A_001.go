package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	if a%3 == 0 || b%3 == 0 || (a+b)%3 == 0 {
		fmt.Println("Possible")
	} else {
		fmt.Println("Impossible")
	}
}
