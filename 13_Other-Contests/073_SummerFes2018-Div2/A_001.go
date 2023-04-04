package main

import "fmt"

func main() {
	var X, Y, Z int
	fmt.Scan(&X, &Y, &Z)
	if X == Y || Y == Z || Z == X {
		fmt.Println(2)
	} else if X+Y == 2*Z || Y+Z == 2*X || Z+X == 2*Y {
		fmt.Println(3)
	} else {
		fmt.Println(-1)
	}
}
