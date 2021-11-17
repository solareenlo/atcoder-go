package main

import "fmt"

func main() {
	var c string
	fmt.Scan(&c)

	if c[0] == c[1] && c[1] == c[2] {
		fmt.Println("Won")
	} else {
		fmt.Println("Lost")
	}
}
