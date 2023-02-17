package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	l := len(s)
	vis := make([]bool, 255)
	cnt := 8
	ans := make([]byte, 9)
	for i := l - 1; i >= 0; i-- {
		if !vis[s[i]] {
			ans[cnt] = s[i]
			cnt--
			vis[s[i]] = true
		}
	}
	if cnt > 0 {
		for i := 'A'; i <= 'H'; i++ {
			if !vis[i] {
				ans[cnt] = byte(i)
				cnt--
			}
		}
	}
	for i := 1; i <= 8; i++ {
		fmt.Printf("%c", ans[i])
	}
	fmt.Println()
}
