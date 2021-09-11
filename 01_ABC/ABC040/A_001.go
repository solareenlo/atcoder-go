package main

import "fmt"

func ma() {
	var n, x int
	fmt.Scan(&n, &x)
	if n/2 >= x {
		fmt.Println(x - 1)
	} else {
		fmt.Println(n - x)
	}
}
