package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	var a string
	fmt.Scan(&a)
	ans := ""
	for i := 0; i < len(s); i++ {
		s1 := ""
		s1 += string(s[i])
		if s[i] == a[0] {
			s1 += a
		}
		ans += s1
	}
	fmt.Println(ans)
}
