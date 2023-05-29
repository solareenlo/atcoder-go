package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	if a != 0 && b != 0 {
		fmt.Println(2)
	} else if a == 0 && b == 0 {
		fmt.Println(0)
	} else {
		fmt.Println(1)
	}
}
