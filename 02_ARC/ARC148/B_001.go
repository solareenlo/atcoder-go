package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	var s string
	fmt.Scan(&n, &s)
	ans := s
	for i := 0; i < len(s); i++ {
		if s[i] == 'p' {
			for j := i; j < len(s); j++ {
				t := strings.Split(s, "")
				reverseOrderString(t, i, j+1)
				for k := i; k <= j; k++ {
					if t[k] == "d" {
						t[k] = "p"
					} else {
						t[k] = "d"
					}
				}
				ans = min(ans, strings.Join(t, ""))
			}
			break
		}
	}
	fmt.Println(ans)
}

func min(a, b string) string {
	if len(a) < len(b) || (len(a) == len(b)) && a < b {
		return a
	}
	return b
}

func reverseOrderString(a []string, start, end int) {
	for i, j := start, end-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}
