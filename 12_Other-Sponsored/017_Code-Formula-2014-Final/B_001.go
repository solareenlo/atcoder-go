package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	x := n / 2
	n %= 2
	if n != 0 {
		fmt.Println(x + 3)
	} else {
		fmt.Println(x)
	}
}
