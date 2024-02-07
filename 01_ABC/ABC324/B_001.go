package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for n%2 == 0 {
		n /= 2
	}
	for n%3 == 0 {
		n /= 3
	}
	if n == 1 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
