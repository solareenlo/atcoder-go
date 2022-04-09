package main

import "fmt"

func main() {
	var n int
	var s string
	fmt.Scan(&n, &s)

	fmt.Println(string(s[n-1]))
}
