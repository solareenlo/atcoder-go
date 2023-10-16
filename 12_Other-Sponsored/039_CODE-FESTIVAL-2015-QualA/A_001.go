package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Print(s[:len(s)-1], "5")
	fmt.Println()
}
