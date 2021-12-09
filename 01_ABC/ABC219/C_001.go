package main

import (
	"fmt"
	"sort"
)

func main() {
	var x []byte
	var n int
	fmt.Scan(&x, &n)

	y := []byte("abcdefghijklmnopqrstuvwxyz")
	m := make(map[byte]byte)
	for i := range x {
		m[x[i]] = y[i]
	}

	s := make([][2]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&s[i][0])
		for j := 0; j < len(s[i][0]); j++ {
			s[i][1] += string(m[s[i][0][j]])
		}
	}
	sort.Slice(s, func(i, j int) bool {
		return s[i][1] < s[j][1]
	})

	for i := range s {
		fmt.Println(s[i][0])
	}
}
