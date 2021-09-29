package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	for i := range s {
		if i%2 == 0 {
			fmt.Print(string(s[i]))
		}
	}
}
