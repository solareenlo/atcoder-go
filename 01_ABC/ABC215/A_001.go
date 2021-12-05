package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	switch s {
	case "Hello,World!":
		fmt.Println("AC")
	default:
		fmt.Println("WA")
	}
}
