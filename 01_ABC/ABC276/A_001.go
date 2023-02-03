package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	ans := -1
	for i := 0; i < len(s); i++ {
		if s[i] == 'a' {
			ans = i + 1
		}
	}
	fmt.Println(ans)
}
