package main

import "fmt"

func main() {
	var s string
	var n int
	fmt.Scan(&s, &n)

	rs := []rune(s)
	var l, r int
	for i := 0; i < n; i++ {
		fmt.Scan(&l, &r)
		for j := 0; j < (r-l)/2+1; j++ {
			rs[l-1+j], rs[r-1-j] = rs[r-1-j], rs[l-1+j]
		}
	}
	fmt.Println(string(rs))
}
