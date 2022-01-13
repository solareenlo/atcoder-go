package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)

	s = strings.ToUpper(s)
	posI := max(strings.Index(s, "I"), 0)
	posC := max(strings.Index(s[posI:], "C"), 0)
	posC += posI
	posT := max(strings.Index(s[posC:], "T"), 0)
	posT += posC

	if posI < posC && posC < posT {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
