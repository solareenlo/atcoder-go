package main

import (
	"fmt"
	"strings"
)

func main() {
	var n, r int
	var s string
	fmt.Scan(&n, &r, &s)
	s += strings.Repeat(" ", r)

	cnt := 0
	for i := 0; i < n; i++ {
		target := false
		for j := i; j < i+r; j++ {
			if s[j] == '.' {
				target = true
			}
		}
		if target {
			for j := i; j < i+r; j++ {
				if j < n {
					tmp := strings.Split(s, "")
					tmp[j] = "o"
					s = strings.Join(tmp, "")
				}
			}
			cnt++
		}
		point := 0
		for j := i; j < n; j++ {
			if s[j] == '.' {
				for k := j; k < j+r; k++ {
					if s[k] == '.' {
						point = k
					}
				}
				if point != 0 {
					point -= r - 1
				}
				break
			}
		}
		if point > i {
			cnt += point - i
			i = point - 1
		}
	}

	fmt.Println(cnt)
}
