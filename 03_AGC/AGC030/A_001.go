package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if a+b >= c {
		fmt.Println(b + c)
	} else {
		fmt.Println(b*2 + a + 1)
	}
}
