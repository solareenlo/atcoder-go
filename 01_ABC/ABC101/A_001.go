package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	cnt := 0
	for i := range s {
		if s[i] == '+' {
			cnt++
		} else {
			cnt--
		}
	}
	fmt.Println(cnt)
}
