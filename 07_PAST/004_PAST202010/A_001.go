package main

import (
	"fmt"
	"sort"
)

func main() {
	v := make([]int, 3)
	m := make(map[int]byte)
	for i := 0; i < 3; i++ {
		fmt.Scan(&v[i])
		m[v[i]] = 'A' + byte(i)
	}
	sort.Ints(v)
	fmt.Println(string(m[v[1]]))
}
