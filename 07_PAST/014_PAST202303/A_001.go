package main

import "fmt"

func main() {
	var n, s, t int
	fmt.Scan(&n, &s, &t)
	if ((n + s + t) % 2) != 0 {
		fmt.Println("Alice")
	} else {
		fmt.Println("Bob")
	}
}
