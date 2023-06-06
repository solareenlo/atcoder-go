package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)

	h := 0
	for i := 0; i < len(s); i++ {
		h += int(s[i]-'a') + 1
	}

	if h == 26*20 || s == "a" {
		fmt.Println("NO")
		return
	}

	ans := make([]string, 0)
	for h > 0 {
		if h > 26 {
			ans = append(ans, "z")
			h -= 26
		} else {
			ans = append(ans, string('a'+h-1))
			h = 0
		}
		if ans[0][0] == s[0] {
			ans[0] = string(ans[0][0] - 1)
			h++
		}
	}
	fmt.Println(strings.Join(ans, ""))
}
