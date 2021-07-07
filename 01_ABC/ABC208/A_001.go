package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	if b < a || b > a*6 {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}
