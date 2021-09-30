package main

import "fmt"

func main() {
	var h, w int
	fmt.Scan(&h, &w)
	s := make([]string, h)
	for i := range s {
		fmt.Scan(&s[i])
		s[i] = "." + s[i] + "."
	}
	t := make([]byte, w+2)
	for i := range t {
		t[i] = '.'
	}
	s = append([]string{string(t)}, s...)
	s = append(s, string(t))

	for i := 1; i <= h; i++ {
		for j := 1; j <= w; j++ {
			if s[i][j] == '#' {
				fmt.Print("#")
				continue
			}
			cnt := 0
			for x := i - 1; x <= i+1; x++ {
				for y := j - 1; y <= j+1; y++ {
					if s[x][y] == '#' {
						cnt++
					}
				}
			}
			fmt.Print(cnt)
		}
		fmt.Println()
	}
}
