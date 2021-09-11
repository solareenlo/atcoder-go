package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	if s[len(s)-1] == 'T' {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
