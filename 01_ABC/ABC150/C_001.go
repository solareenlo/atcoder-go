package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	v := make([]int, 0)
	for i := 1; i < n+1; i++ {
		v = append(v, i)
	}

	p := make([]int, n)
	for i := range p {
		fmt.Scan(&p[i])
	}

	q := make([]int, n)
	for i := range q {
		fmt.Scan(&q[i])
	}

	a, b := false, false
	cntP, cntQ, cnt := 0, 0, 0
	for !a || !b {
		if reflect.DeepEqual(v, p) {
			a = true
			cntP = cnt
		}
		if reflect.DeepEqual(v, q) {
			b = true
			cntQ = cnt
		}
		nextPermutation(sort.IntSlice(v))
		cnt++
	}

	fmt.Println(abs(cntP - cntQ))
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
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
