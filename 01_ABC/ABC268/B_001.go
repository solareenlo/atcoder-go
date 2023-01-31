package main

import "fmt"

func main() {
	var s, t string
	fmt.Scan(&s, &t)

	if len(s) > len(t) {
		fmt.Println("No")
	} else {
		if t[:len(s)] == s {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
