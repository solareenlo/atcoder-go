package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if a == b && a != c {
		fmt.Println("Yes")
	} else if b == c && b != a {
		fmt.Println("Yes")
	} else if c == a && c != b {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
