package main

import "fmt"

func main() {

	str := map[string]int{}

	var n int
	fmt.Scan(&n)
	var s string
	for i := 0; i < n; i++ {
		fmt.Scan(&s)
		str[s]++
	}

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&s)
		str[s]--
	}

	maxi := 0
	for _, v := range str {
		maxi = max(maxi, v)
	}
	fmt.Println(max(0, maxi))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
