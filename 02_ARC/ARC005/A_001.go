package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	cnt := 0
	for i := 0; i < n; i++ {
		var w string
		fmt.Scan(&w)
		if w[len(w)-1] == '.' {
			w = w[:len(w)-1]
		}
		if w == "TAKAHASHIKUN" {
			cnt++
		}
		if w == "Takahashikun" {
			cnt++
		}
		if w == "takahashikun" {
			cnt++
		}
	}

	fmt.Println(cnt)
}
