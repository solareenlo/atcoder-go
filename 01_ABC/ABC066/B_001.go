package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	cnt := 0
	for {
		s = s[:len(s)-2]
		cnt++
		if len(s)%2 != 0 {
			continue
		}
		m := len(s) / 2
		if s[:m] == s[m:] {
			break
		}
	}
	fmt.Println(len(s))
}
