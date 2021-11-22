package main

import (
	"fmt"
	"sort"
)

func f(s string, b, c []int) int {
	if b[c[s[0]]] == 0 {
		return -1
	}
	res := 0
	for _, x := range s {
		res = res*10 + b[c[x]]
	}
	return res
}

func main() {
	var s1, s2, s3 string
	fmt.Scan(&s1, &s2, &s3)

	m := map[byte]bool{}
	for i := range s1 {
		m[s1[i]] = true
	}
	for i := range s2 {
		m[s2[i]] = true
	}
	for i := range s3 {
		m[s3[i]] = true
	}
	if len(m) > 10 {
		fmt.Println("UNSOLVABLE")
		return
	}
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)

	c := make([]int, 128)
	n := 0
	for _, x := range keys {
		c[x] = n
		n++
	}

	b := make([]int, 10)
	for i := 0; i < 10; i++ {
		b[i] = i
	}

	x := f(s1, b, c)
	y := f(s2, b, c)
	z := f(s3, b, c)
	if x >= 0 && y >= 0 && z >= 0 && x+y == z {
		fmt.Println(x)
		fmt.Println(y)
		fmt.Println(z)
		return
	}
	for nextPermutation(sort.IntSlice(b)) {
		x := f(s1, b, c)
		y := f(s2, b, c)
		z := f(s3, b, c)
		if x >= 0 && y >= 0 && z >= 0 && x+y == z {
			fmt.Println(x)
			fmt.Println(y)
			fmt.Println(z)
			return
		}
	}
	fmt.Println("UNSOLVABLE")
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
