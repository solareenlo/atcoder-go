package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n%1111 == 0 {
		fmt.Println("SAME")
	} else {
		fmt.Println("DIFFERENT")
	}
}
