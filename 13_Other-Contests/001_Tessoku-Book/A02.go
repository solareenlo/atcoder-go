package main

import "fmt"

func main() {
	var n, x int
	fmt.Scan(&n, &x)

	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		if a == x {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}
