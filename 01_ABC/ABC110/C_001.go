package main

import "fmt"

func main() {
	var s, t string
	fmt.Scan(&s, &t)

	cntS, cntT := make([]byte, 26), make([]byte, 26)
	n := len(s)
	for i := 0; i < n; i++ {
		cntS[s[i]-'a']++
		cntT[t[i]-'a']++
		if cntS[s[i]-'a'] != cntT[t[i]-'a'] {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
