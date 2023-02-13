package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	n := len(s)
	ans := 0
	for i := 0; i < 1<<n; i++ {
		var t string
		for j := 0; j < n; j++ {
			if ((i >> j) & 1) != 0 {
				t += string(s[j])
			}
		}
		flag := len(t) > 2 && len(t)%2 != 0 && t[len(t)/2-1:len(t)/2-1+3] == "iwi"
		for j := 0; j < len(t); j++ {
			if flag && check(t[j], t[len(t)-j-1]) {
				flag = true
			} else {
				flag = false
			}
		}
		if flag {
			ans = max(ans, len(t))
		}
	}
	fmt.Println(ans)
}

func check(l, r byte) bool {
	return l == 'i' && r == 'i' || l == 'w' && r == 'w' || l == '(' && r == ')' || l == ')' && r == '(' || l == '[' && r == ']' || l == ']' && r == '[' || l == '{' && r == '}' || l == '}' && r == '{'
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
