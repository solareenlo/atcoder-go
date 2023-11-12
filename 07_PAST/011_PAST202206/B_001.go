package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	var a [26][26]int
	for i := 0; i < len(s)-1; i++ {
		a[s[i]-'a'][s[i+1]-'a']++
	}

	t, u, v := 0, 0, 0
	for i := 0; i < 26; i++ {
		for j := 0; j < 26; j++ {
			if t < a[i][j] {
				u = i
				v = j
				t = a[i][j]
			}
		}
	}
	fmt.Println(string(u+97) + string(v+97))
}
