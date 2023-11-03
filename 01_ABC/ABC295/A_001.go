package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		var s string
		fmt.Scan(&s)
		if s == "and" || s == "not" || s == "that" || s == "the" || s == "you" {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}
