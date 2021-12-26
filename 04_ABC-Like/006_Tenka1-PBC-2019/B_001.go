package main

import "fmt"

func main() {
	var n, k int
	var s string
	fmt.Scan(&n, &s, &k)
	k--

	for i := 0; i < n; i++ {
		if s[i] == s[k] {
			fmt.Print(string(s[k]))
		} else {
			fmt.Print("*")
		}
	}
	fmt.Println()
}
