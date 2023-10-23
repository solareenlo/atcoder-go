package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	c := 0
	a := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '2' {
			c++
			a = max(a, c)
		} else {
			c--
		}
		if c < 0 {
			fmt.Println(-1)
			return
		}
	}
	if c != 0 {
		fmt.Println(-1)
	} else {
		fmt.Println(a)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
