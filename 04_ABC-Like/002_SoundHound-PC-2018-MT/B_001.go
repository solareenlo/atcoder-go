package main

import "fmt"

func main() {
	var s string
	var n int
	fmt.Scan(&s, &n)

	for i := 0; i < len(s); i += n {
		fmt.Print(string(s[i]))
	}
	fmt.Println()
}
