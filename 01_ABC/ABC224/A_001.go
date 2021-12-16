package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	if s[len(s)-2:] == "er" {
		fmt.Println("er")
	} else {
		fmt.Println("ist")
	}
}
