package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	if x != 0 && x%100 == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
