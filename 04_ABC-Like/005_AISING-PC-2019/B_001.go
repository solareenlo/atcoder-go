package main

import "fmt"

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)

	p := make([]int, n)
	for i := range p {
		fmt.Scan(&p[i])
	}

	cnt1, cnt2, cnt3 := 0, 0, 0
	for i := 0; i < n; i++ {
		if p[i] <= a {
			cnt1++
		} else if p[i] <= b {
			cnt2++
		} else {
			cnt3++
		}
	}

	fmt.Println(min(cnt1, cnt2, cnt3))
}

func min(a ...int) int {
	res := a[0]
	for i := range a {
		if res > a[i] {
			res = a[i]
		}
	}
	return res
}
