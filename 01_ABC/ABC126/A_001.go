package main

import (
	"fmt"
	"strings"
)

func main() {
	var n, k int
	var s string
	fmt.Scan(&n, &k, &s)

	for i := 0; i < n; i++ {
		if i == k-1 {
			fmt.Print(strings.ToLower(string(s[i])))
		} else {
			fmt.Print(string(s[i]))
		}
	}
}
