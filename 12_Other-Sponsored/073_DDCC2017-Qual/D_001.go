package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var h, w, a, b int
	fmt.Fscan(in, &h, &w, &a, &b)
	s := make([]string, h)
	for i := range s {
		fmt.Fscan(in, &s[i])
	}
	t, y, r := 0, 0, 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if s[i][j] == '.' {
				continue
			}
			if s[h-1-i][j] == 'S' {
				t++
			}
			if s[i][w-1-j] == 'S' {
				y++
			}
			if s[h-1-i][j] == 'S' && s[i][w-1-j] == 'S' && s[h-1-i][w-1-j] == 'S' {
				r++
			}
		}
	}
	t /= 2
	y /= 2
	r /= 4
	ans := a + b
	ans += r * (max(a, b)*2 + min(a, b))
	t -= r * 2
	y -= r * 2
	ans += max(t*a, y*b)
	T, Y := true, true
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if s[i][j] != s[h-1-i][j] {
				T = false
			}
			if s[i][j] != s[i][w-1-j] {
				Y = false
			}
		}
	}
	if T {
		ans -= a
	}
	if Y {
		ans -= b
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
