package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	a = a*3600 + b*60
	c = c*3600 + d*60 + 1
	if a < c {
		fmt.Println("Takahashi")
	} else {
		fmt.Println("Aoki")
	}
}
