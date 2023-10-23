package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	t := ""
	st := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '+' || s[i] == '-' || s[i] == '*' || s[i] == '/' {
			st = append(st, s[i])
			continue
		}
		if s[i] == ',' {
			t += string(st[len(st)-1])
			continue
		}
		if s[i] == ')' {
			st = st[:len(st)-1]
		}
		t += string(s[i])
	}
	fmt.Println(t)
}
