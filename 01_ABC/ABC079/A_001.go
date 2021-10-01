package main

import "fmt"

func main() {
	var n string
	fmt.Scan(&n)

	if n[0] == n[1] && n[1] == n[2] {
		fmt.Println("Yes")
	} else if n[1] == n[2] && n[2] == n[3] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
