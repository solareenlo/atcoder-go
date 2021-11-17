package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	type pair struct{ a, b int }
	ab := make([]pair, m)
	for i := range ab {
		fmt.Scan(&ab[i].a, &ab[i].b)
	}

	var k int
	fmt.Scan(&k)
	cd := make([]pair, k)
	for i := range cd {
		fmt.Scan(&cd[i].a, &cd[i].b)
	}

	res := 0
	for bit := 0; bit < 1<<k; bit++ {
		ball := make([]bool, 101)
		for i := 0; i < k; i++ {
			if bit&(1<<i) != 0 {
				ball[cd[i].a] = true
			} else {
				ball[cd[i].b] = true
			}
		}
		cnt := 0
		for i := range ab {
			if ball[ab[i].a] && ball[ab[i].b] {
				cnt++
			}
		}
		res = max(res, cnt)
	}

	fmt.Println(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
