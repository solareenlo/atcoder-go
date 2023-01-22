package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Printf("%d%d\n", n/10%10, n%10)
}
