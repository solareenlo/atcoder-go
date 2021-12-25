package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	n := len(s)

	cnt := 0
	for i := 0; i < n; i++ {
		if s[i] == '2' {
			cnt++
		}
	}
	fmt.Println(cnt)
}
