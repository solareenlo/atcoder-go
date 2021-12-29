package main

import "fmt"

func main() {
	var s, t string
	fmt.Scan(&s, &t)

	res := t[0]
	if s == "Y" {
		res = t[0] - ('a' - 'A')
	}
	fmt.Println(string(res))
}
