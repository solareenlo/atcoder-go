package main

import "fmt"

func main() {
	var n int
	var s string
	fmt.Scan(&n, &s)

	res := n / 2
	if n%2 == 0 {
		res = -1
	} else {
		for i := 0; i < n; i++ {
			if s[i] != 'a'+byte((i+n)%3) {
				res = -1
			}
		}
	}
	fmt.Println(res)
}
