package main

import "fmt"

func main() {
	var A, B, C int
	fmt.Scan(&A, &B, &C)

	if A == B && B == C && A%2 == 0 {
		fmt.Println(-1)
	} else {
		x := 0
		for A%2 == 0 && B%2 == 0 && C%2 == 0 {
			x++
			a := A
			b := B
			c := C
			A = (b + c) / 2
			B = (c + a) / 2
			C = (a + b) / 2
		}
		fmt.Println(x)
	}
}
