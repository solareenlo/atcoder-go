package main

import "fmt"

func f(rest int, s string) {
	if rest == 0 {
		fmt.Println(s)
	} else {
		for c := 'a'; c < 'd'; c++ {
			f(rest-1, s+string(c))
		}
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	f(n, "")
}
