package main

import "fmt"

func main() {
	var n int
	var s string
	fmt.Scan(&n, &s)
	k := 0
	c := "0"
	ans := 0
	var t [30]int
	for i := 0; i < n; i++ {
		if c[0] != s[i] {
			c = string(s[i])
			k = 0
		}
		k++
		if t[c[0]-'a'] < k {
			t[c[0]-'a']++
			ans++
		}
	}
	fmt.Println(ans)
}
