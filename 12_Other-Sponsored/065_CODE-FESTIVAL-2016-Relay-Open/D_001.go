package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	fmt.Println(a, b, 3*c-a-b)
	fmt.Println(4*c-2*a-b, c, 2*a+b-2*c)
	fmt.Println(a+b-c, 2*c-b, 2*c-a)
}
