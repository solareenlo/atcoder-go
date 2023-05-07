package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	p := 1
	for b > 0 {
		b--
		p *= c
	}
	if a < p {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
