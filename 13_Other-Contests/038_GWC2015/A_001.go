package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{25, 39, 51, 76, 163, 111, 136, 128, 133, 138}
	ans := make([]int, 1)
	for i := 0; i < len(a); i++ {
		pre := len(ans)
		for j := pre - 1; j >= 0; j-- {
			ans = append(ans, ans[j]+a[i])
		}
		if i == 6 {
			for j := pre - 1; j >= 0; j-- {
				ans = append(ans, ans[j]+58)
			}
		}
	}
	sort.Ints(ans)
	ans = unique(ans)
	for i := range ans {
		fmt.Println(ans[i])
	}
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
	sort.Ints(result)
	n := len(result)
	for i := 0; i < n; i++ {
		a[i] = result[i]
	}
	return result
}
