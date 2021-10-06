package main

import "fmt"

func main() {
	var a, b, x int
	fmt.Scan(&a, &b, &x)

	if x-a < 0 {
		fmt.Println("NO")
	} else if x-a <= b {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
