package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	for i := range s {
		if s[i] == '7' {
			fmt.Println("Yes")
			return
		}
	}

	fmt.Println("No")
}
