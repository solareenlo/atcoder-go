package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	if y < abs(x)*2 || (y-abs(x)*2)%4 != 0 {
		fmt.Println(-1)
	} else {
		fmt.Println(y / 2)
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
