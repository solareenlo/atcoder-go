package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)

	t := strings.Split(s, "+")
	res := 0
	for i := range t {
		if !strings.Contains(t[i], "0") {
			res++
		}
	}
	fmt.Println(res)
}
