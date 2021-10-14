package main

import "fmt"

func main() {
	var s int
	fmt.Scan(&s)

	cnt := 0
	if s == 1 || s == 2 {
		cnt = 0
	} else {
		for s != 4 {
			if s%2 != 0 {
				s = 3*s + 1
			} else {
				s /= 2
			}
			cnt++
		}
	}
	fmt.Println(cnt + 4)
}
