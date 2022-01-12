package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'O':
			fmt.Print(0)
		case 'D':
			fmt.Print(0)
		case 'I':
			fmt.Print(1)
		case 'Z':
			fmt.Print(2)
		case 'S':
			fmt.Print(5)
		case 'B':
			fmt.Print(8)
		default:
			fmt.Print(string(s[i]))
		}
	}
	fmt.Println()
}
