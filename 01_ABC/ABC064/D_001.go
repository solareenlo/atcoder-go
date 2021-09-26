package main

import "fmt"

func main() {
	var n int
	var s string
	fmt.Scan(&n, &s)

	cnt, res := 0, s
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			cnt++
		} else {
			cnt--
			if cnt < 0 {
				res = "(" + res
				cnt++
			}
		}
	}
	for cnt > 0 {
		res += ")"
		cnt--
	}
	fmt.Println(res)
}
