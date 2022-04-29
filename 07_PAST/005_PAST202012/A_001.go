package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	ans := "draw"
	for i := 0; i <= 2; i++ {
		if s[i] == s[i+1] && s[i] == s[i+2] {
			ans = string(s[i])
		}
	}
	fmt.Println(ans)
}
