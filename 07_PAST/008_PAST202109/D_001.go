package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	dx := 0
	dy := 0
	for i := 1; i <= 1000000; i++ {
		if x%i == 0 {
			dx++
		}
		if y%i == 0 {
			dy++
		}
	}

	if dx > dy {
		fmt.Println("X")
	}
	if dx < dy {
		fmt.Println("Y")
	}
	if dx == dy {
		fmt.Println("Z")
	}
}
