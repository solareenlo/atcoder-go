package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n == 1 {
		fmt.Println(1)
	} else {
		fmt.Println(n * 7 / 2)
	}
}
