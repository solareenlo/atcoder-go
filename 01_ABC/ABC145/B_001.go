package main

import "fmt"

func main() {
	var n int
	var s string
	fmt.Scan(&n, &s)

	if n%2 != 0 {
		fmt.Println("No")
	} else {
		sub1 := s[:n/2]
		sub2 := s[n/2:]
		if sub1 == sub2 {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
