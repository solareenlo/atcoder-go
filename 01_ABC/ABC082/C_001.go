package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	m := map[int]int{}
	var a int
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		m[a]++
	}

	res := 0
	for k, v := range m {
		if k > v {
			res += v
		} else {
			res += abs(k - v)
		}
	}
	fmt.Println(res)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
