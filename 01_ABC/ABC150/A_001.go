package main

import "fmt"

func main() {
	var k, x int
	fmt.Scan(&k, &x)

	if 500*k >= x {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
