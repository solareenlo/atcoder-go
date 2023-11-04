package main

import "fmt"

func main() {
	var n int
	var s string
	fmt.Scan(&n, &s)
	c := 0
	for _, x := range s {
		if x == '|' {
			c++
		} else if x == '*' {
			break
		}
	}
	if c == 1 {
		fmt.Println("in")
	} else {
		fmt.Println("out")
	}
}
