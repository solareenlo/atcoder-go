package main

import "fmt"

func main() {
	var s string
	var n int
	fmt.Scan(&s, &n)
	n--
	fmt.Printf("%c%c\n", s[n/5], s[n%5])
}
