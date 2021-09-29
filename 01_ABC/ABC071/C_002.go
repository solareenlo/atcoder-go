package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}

	v := make([]int, 0)
	cnt := map[int]int{}
	for i := 0; i < n; i++ {
		cnt[a[i]]++
		if cnt[a[i]] == 2 {
			cnt[a[i]] = 0
			v = append(v, a[i])
		}
	}
	sort.Ints(v)

	res := 0
	if 2 <= len(v) {
		res = v[len(v)-1] * v[len(v)-2]
	}
	fmt.Println(res)
}
