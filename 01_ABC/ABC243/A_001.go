package main

import "fmt"

func main() {
	var v, a, b, c int
	fmt.Scan(&v, &a, &b, &c)

	v = v % (a + b + c)
	if a > v {
		fmt.Println("F")
	} else if a+b > v {
		fmt.Println("M")
	} else {
		fmt.Println("T")
	}
}
