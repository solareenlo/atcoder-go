package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	rem := n % 1000
	if rem == 0 {
		fmt.Println(0)
	} else {
		fmt.Println(1000 - rem)
	}
}
