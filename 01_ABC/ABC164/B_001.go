package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	for {
		c -= b
		if c <= 0 {
			fmt.Println("Yes")
			break
		}
		a -= d
		if a <= 0 {
			fmt.Println("No")
			break
		}
	}
}
