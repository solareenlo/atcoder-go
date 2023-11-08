package main

import "fmt"

func main() {
	var s [105]string
	var S [4]string

	var n, m int
	fmt.Scan(&n, &m)
	S[0] = "###."
	S[1] = "###."
	S[2] = "###."
	S[3] = "...."
	for i := 1; i <= n; i++ {
		fmt.Scan(&s[i])
		s[i] = "-" + s[i]
	}
	for i := 1; i+8 <= n; i++ {
		for j := 1; j+8 <= m; j++ {
			f := true
			for o := 0; o < 4; o++ {
				for p := 0; p < 4; p++ {
					if s[i+o][j+p] != S[o][p] || s[i+8-o][j+8-p] != S[o][p] {
						f = false
					}
				}
			}
			if f {
				fmt.Println(i, j)
			}
		}
	}
}
