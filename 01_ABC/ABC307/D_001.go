package main

import "fmt"

func main() {
	var n int
	var s string
	fmt.Scan(&n, &s)
	top, s1 := 0, 0
	var st [500500]byte
	for i := 0; i < n; i++ {
		top++
		st[top] = s[i]
		if s[i] == '(' {
			s1++
		}
		if s[i] == ')' {
			if s1 != 0 {
				s1--
				for st[top] != '(' {
					top--
				}
				top--
			}
		}
	}
	for i := 1; i <= top; i++ {
		fmt.Print(string(st[i]))
	}
	fmt.Println()
}
