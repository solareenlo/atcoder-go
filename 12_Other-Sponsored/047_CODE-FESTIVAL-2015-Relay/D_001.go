package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n == 1 {
		fmt.Println(n)
	} else {
		fmt.Println(2 * n)
	}
}
