package main

import "fmt"

func main() {
	var A, B int
	var op string
	fmt.Scan(&A, &op, &B)

	switch op {
	case "+":
		fmt.Println(A + B)
	case "-":
		fmt.Println(A - B)
	case "*":
		fmt.Println(A * B)
	case "/":
		if B != 0 {
			fmt.Println(A / B)
		} else {
			fmt.Println("error")
		}
	default:
		fmt.Println("error")
	}
}
