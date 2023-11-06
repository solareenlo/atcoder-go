package main

import "fmt"

func main() {
	const N = 300300
	const INF = int(1e18)

	var f [N][10]int

	var x, y, z int
	var s string
	fmt.Scan(&x, &y, &z)
	fmt.Scan(&s)
	f[0][1] = INF
	f[0][0] = 0
	s = "!" + s
	for i := 1; i < len(s); i++ {
		if s[i] == 'a' {
			f[i][0] = min(f[i-1][0]+x, f[i-1][1]+min(x, y)+z)
			f[i][1] = min(f[i-1][1]+y, f[i-1][0]+min(x, y)+z)
		} else {
			f[i][0] = min(f[i-1][0]+y, f[i-1][1]+min(x, y)+z)
			f[i][1] = min(f[i-1][1]+x, f[i-1][0]+min(x, y)+z)
		}
	}
	fmt.Println(min(f[len(s)-1][0], f[len(s)-1][1]))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
