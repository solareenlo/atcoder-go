package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	if 3*x >= y && 3*y >= x && (3*x-y)%8 == 0 && (3*y-x)%8 == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
