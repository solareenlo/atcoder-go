package main

import "fmt"

func main() {
	a := []int{-1, 0, 1, 0, 2, 1, 2, 1, 1, 2, 1, 2, 1}

	var x, y int
	fmt.Scan(&x, &y)
	if a[x] == a[y] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
