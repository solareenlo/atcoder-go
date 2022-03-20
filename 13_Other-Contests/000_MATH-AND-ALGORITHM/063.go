package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n&1 != 0 {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}
