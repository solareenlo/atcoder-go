package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	l, r := -1, -1
	for i := 1; i < len(s); i++ {
		if s[i-1] == s[i] {
			l, r = i, i+1
		}
	}
	for i := 2; i < len(s); i++ {
		if s[i-2] == s[i] {
			l, r = i-1, i+1
		}
	}
	fmt.Println(l, r)
}
