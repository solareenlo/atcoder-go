package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	if s[0] == '1' || s == "0000000000" {
		fmt.Println("No")
	} else {
		if (s[1] == '0' && s[7] == '0') || (s[2] == '0' && s[8] == '0') || (s[0] == '0' && s[4] == '0') {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
