package main

import "fmt"

func main() {
	var h, w int
	fmt.Scan(&h, &w)

	var x, y [2]int
	k := 0
	for i := 0; i < h; i++ {
		var s string
		fmt.Scan(&s)
		for j := 0; j < w; j++ {
			if string(s[j]) == "o" {
				x[k] = i
				y[k] = j
				k++
			}
		}
	}

	fmt.Println(abs(x[0]-x[1]) + abs(y[0]-y[1]))
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
