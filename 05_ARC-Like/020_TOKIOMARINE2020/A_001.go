package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	for i := 0; i < 3; i++ {
		fmt.Print(string(s[i]))
	}
}
