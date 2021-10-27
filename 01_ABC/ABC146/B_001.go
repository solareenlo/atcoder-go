package main

import "fmt"

func main() {
	var n int
	var s string
	fmt.Scan(&n, &s)

	for i := range s {
		if s[i]+byte(n) <= 'Z' {
			fmt.Print(string(s[i] + byte(n)))
		} else {
			fmt.Print(string(s[i] + byte(n-26)))
		}
	}
	fmt.Println()
}
