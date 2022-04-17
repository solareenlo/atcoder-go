package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	for i := range s {
		if s[i] == 'C' {
			for j := i + 1; j < len(s); j++ {
				if s[j] == 'F' {
					fmt.Println("Yes")
					return
				}
			}
		}
	}
	fmt.Println("No")
}
