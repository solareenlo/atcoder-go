package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)

	num := len(s)
	for i := 0; i < num; i++ {
		if s[i] == '*' {
			tmp := strings.Split(s, "")
			tmp[i] = string(s[num-1-i])
			s = strings.Join(tmp, "")
		}
	}

	for i := 0; i < num; i++ {
		if s[i] == '*' {
			tmp := strings.Split(s, "")
			tmp[i] = "a"
			s = strings.Join(tmp, "")
		}
	}

	t := reverseString(s)

	if s == t {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func reverseString(s string) string {
	res := []rune(s)
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}
