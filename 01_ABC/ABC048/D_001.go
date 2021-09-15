package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	n := len(s)
	res := "First"
	if (s[0] == s[n-1]) != (n%2 == 0) {
		res = "Second"
	}
	fmt.Println(res)
}
