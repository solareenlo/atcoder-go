package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	n := len(s)
	cnt := 0
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-1-i] {
			cnt++
		}
	}
	fmt.Println(cnt)
}
