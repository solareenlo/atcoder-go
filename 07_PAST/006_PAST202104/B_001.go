package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	for i := 1; i < len(s); i += 4 {
		if s[i] == 'o' {
			fmt.Println((i + 3) / 4)
			return
		}
	}
	fmt.Println("none")
}
