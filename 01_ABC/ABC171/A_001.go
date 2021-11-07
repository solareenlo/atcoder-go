package main

import "fmt"

func main() {
	var c string
	fmt.Scan(&c)

	if "A" <= c && c <= "Z" {
		fmt.Println("A")
	} else {
		fmt.Println("a")
	}
}
