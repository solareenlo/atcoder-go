package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	if a == b%10+1 || b == a%10+1 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
