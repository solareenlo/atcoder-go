package main

import "fmt"

func main() {
	var A, B string
	fmt.Scan(&A, &B)

	a, b := 0, 0
	for i := range A {
		a += int(A[i] - '0')
	}
	for i := range B {
		b += int(B[i] - '0')
	}

	if a < b {
		fmt.Println(b)
	} else {
		fmt.Println(a)
	}
}
