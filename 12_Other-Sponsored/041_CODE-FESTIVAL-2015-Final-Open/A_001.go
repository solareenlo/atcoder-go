package main

import (
	"fmt"
)

func main() {
	var a, b, c string
	fmt.Scan(&a, &b, &c)
	if len(a) == 5 && len(b) == 7 && len(c) == 5 {
		fmt.Println("valid")
	} else {
		fmt.Println("invalid")
	}
}
