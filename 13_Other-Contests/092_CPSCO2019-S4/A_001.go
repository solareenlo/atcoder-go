package main

import "fmt"

func main() {
	var l, x int
	fmt.Scan(&l, &x)
	if (x/l)%2 == 1 {
		fmt.Println(l - x%l)
	} else if x < l {
		fmt.Println(x)
	} else {
		fmt.Println(x % l)
	}
}
