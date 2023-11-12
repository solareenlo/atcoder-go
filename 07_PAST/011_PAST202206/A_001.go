package main

import "fmt"

func main() {
	var X, A, B, C int
	fmt.Scan(&X, &A, &B, &C)
	d := X*B + A*B*C - X*A
	if d > 0 {
		fmt.Println("Tortoise")
	} else if d < 0 {
		fmt.Println("Hare")
	} else {
		fmt.Println("Tie")
	}
}
