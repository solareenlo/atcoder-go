package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n%5 >= 3 {
		fmt.Println(n + (5 - n%5))
	} else {
		fmt.Println(n - n%5)
	}
}
