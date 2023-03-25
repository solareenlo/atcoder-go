package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n == 45 {
		fmt.Println(1)
		return
	}

	x := 90
	if x%n == 0 {
		fmt.Println(x / n)
	} else {
		for x%n != 0 {
			x += 90
		}
		fmt.Println(x / n)
	}
}
