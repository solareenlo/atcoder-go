package main

import "fmt"

func main() {
	var x, y, z int
	fmt.Scan(&x, &y, &z)

	if y*(y-x) > 0 {
		fmt.Println(abs(x))
	} else if y*(y-z) > 0 {
		fmt.Println(abs(z) + abs(x-z))
	} else {
		fmt.Println(-1)
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
