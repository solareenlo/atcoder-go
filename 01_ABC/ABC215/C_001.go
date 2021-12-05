package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var s string
	var k int
	fmt.Scan(&s, &k)

	ss := strings.Split(s, "")
	sort.Strings(ss)

	res := ss
	for i := 1; i < k; i++ {
		if nextPermutation(sort.StringSlice(ss)) {
			res = ss
		}
	}
	fmt.Println(strings.Join(res, ""))
}

func nextPermutation(x sort.Interface) bool {
	n := x.Len() - 1
	if n < 1 {
		return false
	}
	j := n - 1
	for ; !x.Less(j, j+1); j-- {
		if j == 0 {
			return false
		}
	}
	l := n
	for !x.Less(j, l) {
		l--
	}
	x.Swap(j, l)
	for k, l := j+1, n; k < l; {
		x.Swap(k, l)
		k++
		l--
	}
	return true
}
