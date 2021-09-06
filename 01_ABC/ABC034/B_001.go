package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n&1 == 1 {
		fmt.Println(n + 1)
	} else {
		fmt.Println(n - 1)
	}
}
