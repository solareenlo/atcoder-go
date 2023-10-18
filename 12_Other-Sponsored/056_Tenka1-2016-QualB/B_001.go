package main

import "fmt"

func main() {
	var S string
	fmt.Scan(&S)
	ss := len(S)
	l, r, mu := 0, 0, 0
	cnt := 0
	for i := ss - 1; i >= 0; i-- {
		if S[i] == '(' {
			r++
		} else {
			l++
		}
		if l < r {
			if mu == 0 {
				mu = i
			}
			r--
			l++
			cnt++
		}
		if l > ss/2 {
			if mu == 0 {
				mu = i
			}
			l--
			r++
			cnt++
		}
	}
	fmt.Println(cnt + mu)
}
