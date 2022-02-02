package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	n := len(s)
	tmp := 0
	if s[0] == s[n-1] {
		tmp = 1
	}
	if (n&1)^tmp != 0 {
		fmt.Println("First")
	} else {
		fmt.Println("Second")
	}
}
