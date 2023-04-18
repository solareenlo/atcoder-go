package main

import "fmt"

func main() {
	var A, B int
	fmt.Scan(&A, &B)
	if A >= B {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}
