package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	fmt.Println(n + (6-n%4)%4)
}
