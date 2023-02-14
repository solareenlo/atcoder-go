package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	if s[0] == 'x' && s[len(s)-1] == 'x' {
		fmt.Println("x")
	} else {
		fmt.Println("o")
	}
}
