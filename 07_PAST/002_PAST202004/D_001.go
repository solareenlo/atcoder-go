package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)
	n := len(s)
	st := map[string]bool{}
	for i := 0; i < n; i++ {
		for l := 1; l <= 3; l++ {
			if i+l <= n {
				t := s[i : i+l]
				for k := 0; k < 1<<l; k++ {
					u := strings.Split(t, "")
					for j := 0; j < l; j++ {
						if k>>j&1 != 0 {
							u[j] = "*"
						}
					}
					st[strings.Join(u, "")] = true
				}
			}
		}
	}
	fmt.Println(len(st))
}
