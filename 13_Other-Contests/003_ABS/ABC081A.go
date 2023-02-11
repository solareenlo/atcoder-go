package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	cnt := 0
	for i := range s {
		if s[i] == '1' {
			cnt++
		}
	}
	fmt.Println(cnt)
}
