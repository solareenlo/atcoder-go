package main

import "fmt"

func main() {
	var N, A, B int
	fmt.Scan(&N, &A, &B)
	for i := 0; i < N; i++ {
		var C int
		fmt.Scan(&C)
		if A+B == C {
			fmt.Println(i + 1)
			return
		}
	}
}
