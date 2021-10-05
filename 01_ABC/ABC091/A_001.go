package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if a+b >= c {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
