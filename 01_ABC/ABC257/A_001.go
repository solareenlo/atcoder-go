package main

import "fmt"

func main() {
	var n, x int
	fmt.Scan(&n, &x)

	fmt.Printf("%c\n", 'A'+(x-1)/n)
}
