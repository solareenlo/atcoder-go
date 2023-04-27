package main

import "fmt"

func main() {
	var s int
	fmt.Scan(&s)
	ans := 0
	for s != 0 {
		if s%10 > ans {
			ans = s % 10
		}
		s = s / 10
	}
	fmt.Println(ans)
}
