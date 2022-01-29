package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	fmt.Print("1")
	for i := 0; i < n-1; i++ {
		fmt.Print("0")
	}
	fmt.Println("7")
}
