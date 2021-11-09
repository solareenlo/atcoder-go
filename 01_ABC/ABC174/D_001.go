package main

import "fmt"

func main() {
	var n int
	var c string
	fmt.Scan(&n, &c)

	cntW, cntR := 0, 0
	for i := 0; i < n; i++ {
		if c[i] == 'W' {
			cntW++
		}
		if c[i] == 'R' {
			cntR++
		}
	}

	sub := c[n-cntW:]
	cnt := 0
	for i := 0; i < len(sub); i++ {
		if sub[i] == 'R' {
			cnt++
		}
	}

	if cntR == n {
		cnt = 0
	}
	fmt.Println(cnt)
}
