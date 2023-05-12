package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)
	m := make([][]string, 600)
	for i := range m {
		m[i] = make([]string, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < i+1; j++ {
			if i%2 != 0 {
				m[j][i] = "."
			} else {
				m[i][j] = "."
			}
		}
		for j := 0; j < i; j++ {
			if i%2 != 0 {
				m[i][j] = "#"
			} else {
				m[j][i] = "#"
			}
		}
	}
	for i := 0; i < n; i++ {
		fmt.Println(strings.Join(m[i], ""))
	}
}
