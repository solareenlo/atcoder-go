package main

import "fmt"

func main() {
	var x, a, b int
	fmt.Scan(&x, &a, &b)
	x++
	fmt.Println(x)
	fmt.Println(x * (a + b))
	fmt.Println((x * (a + b)) * (x * (a + b)))
	fmt.Println((x*(a+b))*(x*(a+b)) - 1)
}
