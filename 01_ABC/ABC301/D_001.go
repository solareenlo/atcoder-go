package main

import "fmt"

func main() {
	var s string
	var n int
	fmt.Scan(&s, &n)
	a := 1
	var ans [61]int
	res := 0
	for i := len(s) - 1; i >= 0; i, a = i-1, a*2 {
		if s[i] == '?' {
			ans[i] = a
		}
		if s[i] == '1' {
			res += a
		}
	}
	for i := 0; i < len(s); i++ {
		if res+ans[i] <= n {
			res += ans[i]
		}
	}
	if res > n {
		fmt.Println(-1)
	} else {
		fmt.Println(res)
	}
}
