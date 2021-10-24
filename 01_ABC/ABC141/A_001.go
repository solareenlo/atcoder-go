package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	S, C, R := "Sunny", "Cloudy", "Rainy"
	if s == S {
		fmt.Println(C)
	} else if s == C {
		fmt.Println(R)
	} else {
		fmt.Println(S)
	}
}
