package main

import "fmt"

func main() {
	var a, b int
	var c string
	fmt.Scan(&a, &b, &c)
	if len(c) >= a && len(c) <= b {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
