package main

import (
	"fmt"
	"strings"
)

func main() {
	var s, t string
	fmt.Scan(&t, &s)
	for i := len(t) - len(s); i >= 0; i-- {
		r, g := t, true
		for j := range s {
			if !(t[i+j] == '?' || t[i+j] == s[j]) {
				g = false
				break
			}
			r = r[:i+j] + string(s[j]) + r[i+j+1:]
		}
		if g {
			fmt.Println(strings.Replace(r, "?", "a", -1))
			return
		}
	}
	fmt.Println("UNRESTORABLE")
}
