package main

import "fmt"

func main() {
	var n int
	var s1, s2 string
	fmt.Scan(&n, &s1, &s2)

	var res, start int
	if s1[0] == s2[0] {
		res, start = 3, 1
	} else {
		res, start = 3*2, 2
	}

	mod := int(1e9 + 7)
	for i := start; i < n; i++ {
		if s1[i-1] == s2[i-1] {
			res = res * 2 % mod
			if s1[i] != s2[i] {
				i++
			}
		} else if s1[i] != s2[i] {
			res = res * 3 % mod
			i++
		}
	}
	fmt.Println(res)
}
