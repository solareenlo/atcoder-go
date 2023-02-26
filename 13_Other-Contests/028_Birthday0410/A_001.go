package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	if x == 2 {
		fmt.Println("-2147483648 0")
	} else {
		fmt.Println("-1 1")
	}
	fmt.Println("0 0")
}
