package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	l := len(s)
	var b [27]int
	for i := 0; i < l; i++ {
		b[s[i]-'a']++
	}
	ans := 0
	for i := 0; i <= 25; i++ {
		if ans != 0 && b[i] != 0 && i != 0 && ans != b[i] {
			fmt.Println("No")
			return
		} else if b[i] != 0 {
			ans = b[i]
		}
	}
	fmt.Println("Yes")
}
