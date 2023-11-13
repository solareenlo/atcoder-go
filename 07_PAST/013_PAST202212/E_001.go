package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	d := 0
	for _, c := range s {
		if c == '(' {
			d++
		}
		if c == ')' {
			d--
		}
		if d < 0 {
			break
		}
	}
	if d == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
