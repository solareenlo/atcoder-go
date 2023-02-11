package main

import "fmt"

func main() {
	var n int
	var s string
	fmt.Scan(&n, &s)

	if s == "BA" || s[0] == 'A' && s[n-1] == 'B' {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}
