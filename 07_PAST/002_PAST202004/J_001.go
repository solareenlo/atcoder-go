package main

import "fmt"

func dfs(s string) string {
	p := -1
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			p = i
		}
		if s[i] == ')' {
			a := s[:p]
			b := s[p+1 : i]
			c := b
			d := s[i+1:]
			c = reverseString(c)
			return dfs(a + b + c + d)
		}
	}
	return s
}

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(dfs(s))
}

func reverseString(s string) string {
	res := []rune(s)
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}
