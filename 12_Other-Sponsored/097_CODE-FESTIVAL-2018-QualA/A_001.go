package main

import "fmt"

func main() {
	var a, b, c, s int
	fmt.Scan(&a, &b, &c, &s)
	if s >= a+b+c && s <= a+b+c+3 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
