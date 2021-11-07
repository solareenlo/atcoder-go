package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	ok := false
	if (y-2*x)%2 == 0 && (4*x-y)%2 == 0 {
		if y-2*x >= 0 && 4*x-y >= 0 {
			ok = true
		}
	}

	if ok {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
