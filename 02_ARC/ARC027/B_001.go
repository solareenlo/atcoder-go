package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	var n int
	var s, t string
	fmt.Scan(&n, &s, &t)

	for i := 0; i < n; i++ {
		var c, r byte
		if unicode.IsDigit(rune(s[i])) {
			c = t[i]
			r = s[i]
		} else {
			c = s[i]
			r = t[i]
		}
		for j := 0; j < n; j++ {
			if s[j] == c {
				tmp := strings.Split(s, "")
				tmp[j] = string(r)
				s = strings.Join(tmp, "")
			}
			if t[j] == c {
				tmp := strings.Split(t, "")
				tmp[j] = string(r)
				t = strings.Join(tmp, "")
			}
		}
	}

	used := make([]bool, 27)
	res := 1
	for i := 0; i < n; i++ {
		if unicode.IsLetter(rune(s[i])) {
			if !used[s[i]-'A'] {
				used[s[i]-'A'] = true
				if i == 0 {
					res *= 9
				} else {
					res *= 10
				}
			}
		}
	}
	fmt.Println(res)
}
