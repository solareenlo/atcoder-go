package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	n := len(s)
	res := 0
	for i := 0; i < n; i++ {
		if i%2 == 0 && s[i] == 'p' {
			res--
		}
		if i%2 == 1 && s[i] == 'g' {
			res++
		}
	}
	fmt.Println(res)
}
