package main

import "fmt"

var s, t string

func d(x, y int) bool {
	return !(s[x] == t[y] || s[x] == '?' || t[y] == '?')
}

func main() {
	fmt.Scan(&s, &t)
	n := len(s)
	m := len(t)
	cnt := 0
	for i := 0; i < m; i++ {
		if d(n-m+i, i) {
			cnt++
		}
	}
	if cnt == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
	for i := 0; i < m; i++ {
		if d(n-m+i, i) {
			cnt--
		}
		if d(i, i) {
			cnt++
		}
		if cnt == 0 {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
