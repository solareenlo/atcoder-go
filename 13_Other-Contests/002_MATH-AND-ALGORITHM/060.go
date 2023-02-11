package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n%4 != 0 {
		fmt.Println("First")
	} else {
		fmt.Println("Second")
	}
}
