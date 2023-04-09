package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	if a < 123 {
		fmt.Println(a)
	} else {
		fmt.Println(a - 1)
	}
}
