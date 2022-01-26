package main

import "fmt"

func main() {
	var s string
	var a, b, c, d int
	fmt.Scan(&s, &a, &b, &c, &d)

	for i := 0; i < len(s)+1; i++ {
		if i == a || i == b || i == c || i == d {
			fmt.Print("\"")
		}
		if i != len(s) {
			fmt.Print(string(s[i]))
		}
	}
	fmt.Println()
}
