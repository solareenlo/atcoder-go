package main

import "fmt"

func main() {
	var m, n int
	fmt.Scan(&m, &n)
	fmt.Println(m - m/n*(n-1))
}
