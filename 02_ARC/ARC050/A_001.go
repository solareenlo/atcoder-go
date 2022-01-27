package main

import "fmt"

func main() {
	var C, c string
	fmt.Scan(&C, &c)

	if C[0]-'A' == c[0]-'a' {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
