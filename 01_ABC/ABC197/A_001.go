package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	fmt.Printf("%c%c%c\n", s[1], s[2], s[0])
}
