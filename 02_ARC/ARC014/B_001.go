package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var s string
	m := map[string]bool{}
	for i := 0; i < n; i++ {
		var w string
		fmt.Scan(&w)
		if i != 0 {
			if w[0] != s[len(s)-1] || m[w] {
				if i%2 != 0 {
					fmt.Println("WIN")
				} else {
					fmt.Println("LOSE")
				}
				return
			}
		}
		m[w] = true
		s = w
	}
	fmt.Println("DRAW")
}
