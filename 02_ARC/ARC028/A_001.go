package main

import "fmt"

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)

	rem := n % (a + b)
	if rem == 0 {
		fmt.Println("Bug")
	} else if rem <= a {
		fmt.Println("Ant")
	} else {
		fmt.Println("Bug")
	}
}
