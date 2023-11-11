package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var n, k int
	var S string
	fmt.Scan(&n, &k, &S)
	s := strings.Split(S, "")
	cnt := 0
	for i := 0; i < n; i++ {
		if s[i] == "X" {
			cnt++
		}
	}
	if cnt < k {
		for i := 0; i < n; i++ {
			if s[i] == "X" {
				s[i] = "Y"
			} else {
				s[i] = "X"
			}
		}
		k = n - k
	}
	ans := 0
	for i := 1; i < n; i++ {
		if s[i] == "Y" && s[i-1] == "Y" {
			ans++
		}
	}
	cnt = 0
	ok := false
	m := 0
	a := make([]int, 200009)
	for i := 0; i < n; i++ {
		if s[i] == "Y" {
			if ok && cnt != 0 {
				a[m] = cnt
				m++
				cnt = 0
			} else {
				ok = true
				cnt = 0
			}
		} else {
			cnt++
		}
	}
	sort.Ints(a[:m])
	for i := 0; i < m; i++ {
		if a[i] <= k {
			ans += a[i] + 1
			k -= a[i]
		} else {
			break
		}
	}
	if ok {
		fmt.Println(ans + k)
	} else {
		fmt.Println(max(0, k-1))
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
