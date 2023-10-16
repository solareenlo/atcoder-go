package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)

	n := len(s)
	a, b := 0, 0
	flag := false

	for i := 0; i < n; i++ {
		if s[i] == '!' {
			if !flag {
				a = i
			}
			flag = true
			b++
		}
	}

	if !flag {
		a = n
	}

	ans := strings.Repeat("-", a%2)
	if b%2 == 1 {
		ans += "!"
	}
	if b%2 == 0 && b != 0 {
		ans += "!!"
	}

	fmt.Println(ans)
}
