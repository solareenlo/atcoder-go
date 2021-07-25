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

	b := make([]int, n)
	bg := make(map[int]int)
	for i := range b {
		fmt.Scan(&b[i])
		bg[b[i]]++
	}

	set := make(map[int]struct{})
	mp := make(map[int]int)
	for i := 0; i < n; i++ {
		x := a[0] ^ b[i]
		mp = bg
		ok := true
		for j := 0; j < n; j++ {
			if mp[x^a[j]] == 0 {
				ok = false
				break
			}
		}
		if ok == true {
			set[x] = struct{}{}
		}
	}

	fmt.Println(len(set))
	res := make([]int, len(set))
	i := 0
	for k := range set {
		res[i] = k
		i++
	}
	sort.Ints(res)
	for i = 0; i < len(res); i++ {
		fmt.Println(res[i])
	}
}
