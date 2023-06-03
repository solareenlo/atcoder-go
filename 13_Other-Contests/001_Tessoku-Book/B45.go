package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	if a+b+c != 0 {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}
