package main

import "fmt"

func main() {
	var X, Y int64
	fmt.Scanf("%d %d", &X, &Y)

	if X > Y {
		X, Y = Y, X
	}

	if (X == 4 && Y == 4) || (X == 3 && (Y == 1 || (Y >= 8 && Y%4 == 0))) || (X == 2 && Y%4 == 2) || (X == 1 && Y%2 == 1) {
		fmt.Println("rng")
	} else {
		fmt.Println("snuke")
	}
}
