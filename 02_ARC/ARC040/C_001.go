package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	s := make([]string, n+1)
	for i := 0; i < n; i++ {
		fmt.Scan(&s[i])
	}

	res := 0
	for i := 0; i < n; i++ {
		pos := -1
		for j := 0; j < n; j++ {
			if s[i][j] == '.' {
				pos = j
			}
		}
		if pos == -1 {
			continue
		}
		for j := 0; j < pos+1; j++ {
			s[i] = replaceAtIndex(s[i], 'o', j)
		}
		for j := pos; j < n && i < n-1; j++ {
			s[i+1] = replaceAtIndex(s[i+1], 'o', j)
		}
		res++
	}

	fmt.Println(res)
}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}
