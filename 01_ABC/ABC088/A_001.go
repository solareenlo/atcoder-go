package main

import "fmt"

func main() {
	var n, a int
	fmt.Scan(&n, &a)

	if n%500 <= a {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
