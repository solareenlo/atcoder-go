package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)
	stk := []string{string(s[len(s)-1])}
	for i := len(s) - 2; i >= 0; i-- {
		c := string(s[i])
		tmp := c
		for len(stk) > 0 {
			pre := stk[len(stk)-1]
			if tmp <= pre {
				stk = stk[:len(stk)-1]
				tmp += pre
			} else {
				break
			}
		}
		stk = append(stk, tmp)
	}
	if len(stk) == 1 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
