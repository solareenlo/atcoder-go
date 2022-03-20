package main

import "fmt"

func main() {
	var N, X, Y int
	fmt.Scan(&N, &X, &Y)

	if abs(X)+abs(Y) <= N && (X+Y)%2 == N%2 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
