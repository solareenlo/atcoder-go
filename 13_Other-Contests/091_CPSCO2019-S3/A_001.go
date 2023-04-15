package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	for i := 0; i < len(s); i++ {
		if s[i] == 'O' {
			fmt.Print("A")
		} else if s[i] == 'A' {
			fmt.Print("O")
		} else {
			fmt.Print(string(s[i]))
		}
	}
	fmt.Println()
}
