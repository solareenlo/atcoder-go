package main

import "fmt"

func main() {
	var s, t string
	fmt.Scan(&s, &t)

	n := len(s)
	m := len(t)

	var a [26]int
	for i := 0; i < n; i++ {
		a[s[i]-'A']++
	}

	var b [26]int
	for i := 0; i < m; i++ {
		b[t[i]-'A']++
	}

	ans := 0
	for i := 0; i < 26; i++ {
		ans += max(a[i], b[i])
	}

	if b[s[0]-'A'] == 0 || a[t[0]-'A'] == 0 {
		ans++
	}

	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
