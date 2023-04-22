package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)

	v := make([]int, 0)
	for i := 1; i <= n; i++ {
		var s string
		fmt.Fscan(in, &s)
		s = " " + s
		for j := 1; j <= m; j++ {
			if s[j] == '.' {
				v = append(v, 31*i-19*j)
			}
		}
	}
	sort.Ints(v)
	v = unique(v)
	lst := -int(2e9)
	ans := 0
	for _, i := range v {
		if i > lst {
			ans++
			lst = i + 48
		}
	}
	fmt.Println(ans)
}

func unique(a []int) []int {
	occurred := map[int]bool{}
	result := []int{}
	for i := range a {
		if occurred[a[i]] != true {
			occurred[a[i]] = true
			result = append(result, a[i])
		}
	}
	// sort.Ints(result)
	n := len(result)
	for i := 0; i < n; i++ {
		a[i] = result[i]
	}
	return result
}
