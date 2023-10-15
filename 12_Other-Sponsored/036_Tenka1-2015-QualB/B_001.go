package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	Len := len(s)
	if Len == 2 {
		fmt.Println("dict")
		return
	}
	k := 0
	for i := 0; i < Len; i++ {
		if s[i] == '{' {
			k++
		}
		if s[i] == '}' {
			k--
		}
		if k == 1 && s[i] == ':' {
			fmt.Println("dict")
			return
		}
	}
	fmt.Println("set")
}
