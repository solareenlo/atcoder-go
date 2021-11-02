package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	if s[0] == s[1] && s[1] == s[2] {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}
