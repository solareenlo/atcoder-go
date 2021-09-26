package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if (100*a+10*b+c)%4 == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
