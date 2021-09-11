package main

import "fmt"

var dx [9]int = [9]int{0, 1, 0, -1, 0, 1, 1, -1, -1}
var dy [9]int = [9]int{-1, 0, 1, 0, 0, 1, -1, -1, 1}

func main() {
	var h, w int
	fmt.Scan(&h, &w)
	s := make([]string, h)
	for i := 0; i < h; i++ {
		fmt.Scan(&s[i])
	}

	res := make([][]rune, h)
	verify := make([][]rune, h)
	for i := 0; i < h; i++ {
		res[i] = make([]rune, w)
		verify[i] = make([]rune, w)
		for j := 0; j < w; j++ {
			res[i][j] = '#'
			verify[i][j] = '.'
		}
	}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if s[i][j] == '.' {
				for k := 0; k < 9; k++ {
					if i+dx[k] >= 0 && i+dx[k] < h && j+dy[k] >= 0 && j+dy[k] < w {
						res[i+dx[k]][j+dy[k]] = '.'
					}
				}
			}
		}
	}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if res[i][j] == '#' {
				for k := 0; k < 9; k++ {
					if i+dx[k] >= 0 && i+dx[k] < h && j+dy[k] >= 0 && j+dy[k] < w {
						verify[i+dx[k]][j+dy[k]] = '#'
					}
				}
			}
		}
	}

	ok := true
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if s[i][j] != byte(verify[i][j]) {
				ok = false
			}
		}
	}
	if !ok {
		fmt.Println("impossible")
	} else {
		fmt.Println("possible")
		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				fmt.Print(string(res[i][j]))
			}
			fmt.Println()
		}
	}
}
