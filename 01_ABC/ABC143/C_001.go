package main

import "fmt"

func main() {
	var n int
	var s string
	fmt.Scan(&n, &s)
	s += "0"

	cnt := 0
	for i := 0; i < n; i++ {
		if s[i] != s[i+1] {
			cnt++
		}
	}
	fmt.Println(cnt)
}
