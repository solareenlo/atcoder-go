package main

import "fmt"

func main() {
	var s, t string
	fmt.Scan(&s, &t)

	if s[0] == t[1] && s[1] == t[0] && s != "##" {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}
