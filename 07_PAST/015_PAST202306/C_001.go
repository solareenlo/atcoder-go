package main

import "fmt"

func main() {
	var X, Y, R, N int
	fmt.Scan(&X, &Y, &R, &N)
	for i := -N; i <= N; i++ {
		for j := -N; j <= N; j++ {
			x := i - X
			y := j - Y
			if x*x+y*y <= R*R {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
			if j != N {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
