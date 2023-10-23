package main

import "fmt"

func main() {
	var h, m, s, c1, c2 int
	fmt.Scan(&h, &m, &s, &c1, &c2)
	t := 3600*h + 60*m + s
	c1 += (59 * t) / 3600
	c2 += (11 * t) / 43200
	Go := max((3600*c1)/59, (43200*c2)/11) + 1
	Getup := min((3600*(c1+1)-1)/59, (43200*(c2+1)-1)/11)
	if Go > Getup {
		fmt.Println(-1)
	} else {
		Go -= t
		Getup -= t
		fmt.Println(Go, Getup)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
