package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	if a*b*c*d == 252 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
