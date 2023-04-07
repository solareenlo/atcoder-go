package main

import "fmt"

func main() {
	fmt.Println("? 2")
	var S1 string
	fmt.Scan(&S1)

	M := 1
	for i := 29; i >= 1; i-- {
		Q := M + (1 << i)
		if Q >= 1000000000 {
			Q = 1000000000
		}
		fmt.Println("?", Q)
		var S2 string
		fmt.Scan(&S2)
		if S1 != S2 {
			M = Q
		}
	}
	if S1 == "even" {
		fmt.Println("!", M+1)
	} else {
		fmt.Println("!", M)
	}
}
