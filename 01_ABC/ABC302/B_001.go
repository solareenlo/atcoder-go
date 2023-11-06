package main

import "fmt"

func main() {
	var dx [8]int = [8]int{-1, -1, -1, 0, 0, 1, 1, 1}
	var dy [8]int = [8]int{-1, 0, 1, -1, 1, -1, 0, 1}
	var n, m int
	fmt.Scan(&n, &m)
	s := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&s[i])
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			for k := 0; k < 8; k++ {
				str := ""
				for t := 0; t < 5; t++ {
					x := i + t*dx[k]
					y := j + t*dy[k]
					if (x < 0) || (x >= n) || (y < 0) || (y >= m) {
						break
					}
					str += string(s[x][y])
				}
				if str == "snuke" {
					for t := 0; t < 5; t++ {
						x := i + t*dx[k] + 1
						y := j + t*dy[k] + 1
						fmt.Println(x, y)
					}
					return
				}
			}
		}
	}
}
