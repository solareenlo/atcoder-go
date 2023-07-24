package main

import "fmt"

func main() {
	var A string
	fmt.Scan(&A)
	var B int
	fmt.Scan(&B)
	C := len(A)
	D := (B - 1) % C
	fmt.Println(string(A[D]))
}
