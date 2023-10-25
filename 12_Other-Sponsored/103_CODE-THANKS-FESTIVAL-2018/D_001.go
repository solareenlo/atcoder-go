package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	a := "~"
	c := 0
	for i := 0; i < len(s); i++ {
		if a >= string(s[i]) {
			a = string(s[i])
			c++
		}
	}
	fmt.Println(c)
}
