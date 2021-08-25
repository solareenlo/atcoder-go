package main

import "fmt"

func main() {
	var a string
	fmt.Scan(&a)

	switch a {
	case "a":
		fmt.Println(-1)
	default:
		fmt.Println("a")
	}
}
