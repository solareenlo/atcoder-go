package main

import "fmt"

func main() {
	var n int
	var s string
	fmt.Scan(&n, &s)

	if s[n-1] == 'o' {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
