package main

import "fmt"

func main() {
	var S string
	fmt.Scan(&S)
	m := make(map[int]int)
	for _, c := range S {
		m[int(c)]++
	}
	ans, c := 0, 0
	for _, v := range m {
		ans += v / 2 * 2
		c += v % 2
	}
	if c > 0 {
		c--
		ans++
	}
	ans *= ans
	ans += c
	fmt.Println(ans)
}
