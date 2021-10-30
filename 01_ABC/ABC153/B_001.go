package main

import "fmt"

func main() {
	var h, n int
	fmt.Scan(&h, &n)

	sum := 0
	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		sum += a
	}

	if sum >= h {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
