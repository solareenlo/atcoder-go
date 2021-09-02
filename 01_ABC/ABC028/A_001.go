package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	switch {
	case n < 60:
		fmt.Println("Bad")
	case n < 90:
		fmt.Println("Good")
	case n < 100:
		fmt.Println("Great")
	default:
		fmt.Println("Perfect")
	}
}
