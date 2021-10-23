package main

import "fmt"

func main() {
	var s, t string
	fmt.Scan(&s, &t)

	cnt := 0
	for i := 0; i < 3; i++ {
		if s[i] == t[i] {
			cnt++
		}
	}
	fmt.Println(cnt)
}
