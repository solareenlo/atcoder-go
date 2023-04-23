package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	if a < 3 {
		fmt.Println(a*2 + b*2 - 3)
	} else if b < 3 {
		fmt.Println(b*10 + c*5 - 9)
	} else {
		fmt.Println(c/2*21 + 21)
	}
}
