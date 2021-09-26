package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	m := map[byte]struct{}{}
	for i := range s {
		m[s[i]] = struct{}{}
	}
	if len(s) == len(m) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
