package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	if s[2] == s[3] && s[4] == s[5] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
