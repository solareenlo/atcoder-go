package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)
	n := len(s)
	c := make([]int, 30)
	cnt := 1
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			cnt++
		} else if s[i] == ')' {
			for j := 0; j < 26; j++ {
				if c[j] == cnt {
					c[j] = 0
				}
			}
			cnt--
		} else {
			if c[s[i]-'a'] != 0 {
				fmt.Println("No")
				return
			}
			c[s[i]-'a'] = cnt
		}
	}
	fmt.Println("Yes")
}
