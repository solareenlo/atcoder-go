package main

import "fmt"

func main() {
	var x, n int
	fmt.Scan(&x, &n)

	p := make([]int, n)
	for i := range p {
		fmt.Scan(&p[i])
	}

	res := pair{1001, -1}
	for i := -1; i < 102; i++ {
		if find(p, i) == false {
			diff := abs(x - i)
			res = min(res, pair{diff, i})
		}
	}

	fmt.Println(res.y)
}

func find(a []int, v int) bool {
	n := len(a)
	for i := 0; i < n; i++ {
		if a[i] == v {
			return true
		}
	}
	return false
}

type pair struct{ x, y int }

func min(a, b pair) pair {
	if a.x < b.x {
		return a
	}
	if a.x == b.x {
		if a.y < b.y {
			return a
		}
		return b
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
