package main

import "fmt"

func main() {
	var n int
	var s string
	fmt.Scan(&n, &s)

	res := -1
	for i := 0; i < n; i++ {
		s1 := s[:i]
		m1 := map[byte]struct{}{}
		for j := 0; j < len(s1); j++ {
			m1[s1[j]] = struct{}{}
		}

		s2 := s[i:]
		m2 := map[byte]struct{}{}
		for j := 0; j < len(s2); j++ {
			m2[s2[j]] = struct{}{}
		}

		cnt := 0
		for key, _ := range m1 {
			if _, ok := m2[key]; ok {
				cnt++
			}
		}
		res = max(res, cnt)
	}
	fmt.Println(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
