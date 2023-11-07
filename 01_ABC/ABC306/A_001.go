package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var s string
	fmt.Scan(&s)
	for _, c := range s {
		fmt.Print(string(c), string(c))
	}
	fmt.Println()
}
