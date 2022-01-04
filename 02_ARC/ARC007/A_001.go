package main

import "fmt"

func main() {
	var x, s string
	fmt.Scan(&x, &s)

	for i := 0; i < len(s); i++ {
		if s[i] != x[0] {
			fmt.Print(string(s[i]))
		}
	}
	fmt.Println()
}
