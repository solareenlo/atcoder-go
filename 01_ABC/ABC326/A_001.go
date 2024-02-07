package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	if x-y > 3 || x-y < -2 {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}
