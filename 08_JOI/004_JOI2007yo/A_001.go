package main

import "fmt"

func main() {
	var a, b, c, d, e, f, g, h int
	fmt.Scan(&a, &b, &c, &d, &e, &f, &g, &h)
	if a+b+c+d > e+f+g+h {
		fmt.Println(a + b + c + d)
	} else {
		fmt.Println(e + f + g + h)
	}
}
