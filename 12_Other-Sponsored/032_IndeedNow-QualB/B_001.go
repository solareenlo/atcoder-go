package main

import "fmt"

func main() {
	var s, t string
	fmt.Scan(&s, &t)
	n := len(s)
	s += s + s
	ans := -1
	for i := 0; i < n+1; i++ {
		if s[i:i+n] == t {
			ans = n - i
		}
	}
	fmt.Println(ans)
}
