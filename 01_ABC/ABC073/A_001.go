package main

import "fmt"

func main() {
	var n string
	fmt.Scan(&n)

	if n[0] == '9' || n[1] == '9' {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
