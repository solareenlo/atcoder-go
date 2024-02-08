package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	ans := ""
	for _, c := range s {
		ans += string(c)
		if len(ans) >= 3 && ans[len(ans)-3:] == "ABC" {
			ans = ans[:len(ans)-3]
		}
	}
	fmt.Println(ans)
}
