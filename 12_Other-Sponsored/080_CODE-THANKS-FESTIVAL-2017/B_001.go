package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	a := len(s)
	var i, j int
	for i = 0; i < a; i++ {
		for j = 0; j < a-i; j++ {
			if s[i+j] != s[a-1-j] {
				break
			}
		}
		if j == a-i {
			fmt.Println(i)
			return
		}
	}
}
