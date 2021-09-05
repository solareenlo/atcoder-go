package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	var n int
	fmt.Scan(&n)

	m := make(map[string]int)
	for i := 0; i < len(s)-n+1; i++ {
		m[s[i:i+n]]++
	}
	fmt.Println(len(m))
}
