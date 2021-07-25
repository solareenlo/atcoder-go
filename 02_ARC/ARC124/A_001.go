package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	var c string
	var a int
	s := make([]int, 1002)
	p := make([]bool, 1002)
	for i := 1; i <= k; i++ {
		fmt.Scan(&c, &a)
		if c == "L" {
			s[a+1]++
		} else {
			s[1]++
			s[a+1]--
		}
		p[a] = true
	}

	res := 1
	for i := 1; i <= n; i++ {
		s[i] += s[i-1]
		if p[i] == false {
			res = res * s[i] % 998244353
		}
	}
	fmt.Println(res)
}
