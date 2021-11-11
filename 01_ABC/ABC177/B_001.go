package main

import "fmt"

func main() {
	var s, t string
	fmt.Scan(&s, &t)

	res := 0
	for i := 0; i < len(s)-len(t)+1; i++ {
		cnt := 0
		for j := 0; j < len(t); j++ {
			if s[i+j] == t[j] {
				cnt++
			}
		}
		res = max(res, cnt)
	}

	fmt.Println(len(t) - res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
