package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	var n int
	fmt.Scan(&n)
	l := len(s)
	n = n % l
	fmt.Println(s[n:n+l-n] + s[:n])
}
