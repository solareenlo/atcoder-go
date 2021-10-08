package main

import "fmt"

func main() {
	var h, w int
	fmt.Scan(&h, &w)
	s := make([][]bool, h)

	var t string
	for i := 0; i < h; i++ {
		s[i] = make([]bool, w)
		fmt.Scan(&t)
		for j := 0; j < w; j++ {
			if t[j] == '#' {
				s[i][j] = true
			}
		}
	}

	dx := []int{0, 1, 0, -1}
	dy := []int{-1, 0, 1, 0}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if s[i][j] {
				sum := 0
				for k := 0; k < 4; k++ {
					nx := i + dx[k]
					ny := j + dy[k]
					if 0 <= nx && nx < h {
						if 0 <= ny && ny < w {
							if s[nx][ny] {
								sum++
							}
						}
					}
				}
				if sum == 0 {
					fmt.Println("No")
					return
				}
			}
		}
	}
	fmt.Println("Yes")
}
