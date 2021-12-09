package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	switch {
	case x < 40:
		fmt.Println(40 - x)
	case x < 70:
		fmt.Println(70 - x)
	case x < 90:
		fmt.Println(90 - x)
	default:
		fmt.Println("expert")
	}
}
