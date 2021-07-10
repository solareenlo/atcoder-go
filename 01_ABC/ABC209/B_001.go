package main

import "fmt"

func main() {
	var n, x int
	fmt.Scan(&n, &x)

	sum := 0
	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		if i%2 == 1 {
			sum += a - 1
		} else {
			sum += a
		}
	}

	if x >= sum {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
