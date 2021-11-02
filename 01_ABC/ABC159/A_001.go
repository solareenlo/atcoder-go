package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	fmt.Println((n*(n-1) + m*(m-1)) / 2)
}
