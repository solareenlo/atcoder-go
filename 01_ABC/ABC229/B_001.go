package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	for a|b != 0 {
		if a%10+b%10 > 9 {
			fmt.Println("Hard")
			return
		}
		a /= 10
		b /= 10
	}
	fmt.Println("Easy")
}
