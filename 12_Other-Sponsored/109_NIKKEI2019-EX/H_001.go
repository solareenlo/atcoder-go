package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	n = n % 9
	if n%2 == 0 && n != 8 {
		fmt.Println("Lose")
	} else {
		fmt.Println("Win")
	}
}
