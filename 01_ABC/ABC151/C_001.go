package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	ac := map[int]bool{}
	wa := map[int]int{}
	sum := 0
	for i := 0; i < m; i++ {
		var p int
		var s string
		fmt.Scan(&p, &s)
		if s == "AC" && !ac[p] {
			ac[p] = true
			sum += wa[p]
		}
		if s == "WA" {
			wa[p]++
		}
	}

	fmt.Println(len(ac), sum)
}
