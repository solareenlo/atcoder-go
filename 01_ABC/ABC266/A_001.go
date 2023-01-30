package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	fmt.Printf("%c\n", s[len(s)>>1])
}
