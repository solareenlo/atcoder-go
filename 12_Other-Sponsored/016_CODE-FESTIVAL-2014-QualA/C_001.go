package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	a--
	fmt.Println((b/4 - b/100 + b/400) - (a/4 - a/100 + a/400))
}
