package main

import "fmt"

func main() {
	var X, Y, A, B, a, b, c, d int
	fmt.Scan(&X, &Y, &A, &B, &a, &b, &c, &d)
	P := (b*X < A*X+a*(B-A))
	Q := (d*X < A*X+c*(B-A))
	if P != Q {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
