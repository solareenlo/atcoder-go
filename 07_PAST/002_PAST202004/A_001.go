package main

import "fmt"

func main() {
	var s, t string
	fmt.Scan(&s, &t)

	S := int(s[0] - '0')
	if s[0] == 'B' {
		S = -1*int(s[1]-'0') + 1
	}

	T := int(t[0] - '0')
	if t[0] == 'B' {
		T = -1*int(t[1]-'0') + 1
	}
	fmt.Println(abs(S - T))
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
