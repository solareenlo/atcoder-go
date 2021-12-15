package main

import "fmt"

func solve2(X, Y, S, T int) bool {
	for i := 0; i < 2; i++ {
		len := (S + X - 1) / X
		if len < Y && X*(Y-len) >= T {
			return true
		}
		X, Y = Y, X
	}
	return false
}

func solve3(X, Y, A, B, C int) bool {
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			len := (A + X - 1) / X
			if len < Y && solve2(X, Y-len, B, C) {
				return true
			}
			A, B = B, A
			B, C = C, B
		}
		X, Y = Y, X
	}
	return false
}

func main() {
	var X, Y, A, B, C int
	fmt.Scan(&X, &Y, &A, &B, &C)
	if solve3(X, Y, A, B, C) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
