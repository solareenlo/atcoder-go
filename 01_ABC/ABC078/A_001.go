package main

import "fmt"

func main() {
	var x, y string
	fmt.Scan(&x, &y)

	if x < y {
		fmt.Println("<")
	} else if x > y {
		fmt.Println(">")
	} else {
		fmt.Println("=")
	}
}
