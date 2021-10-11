package main

import "fmt"

func main() {
	var h, w int
	fmt.Scan(&h, &w)

	a := [100][100]byte{}
	var s string
	for i := 0; i < h; i++ {
		fmt.Scan(&s)
		for j := 0; j < w; j++ {
			a[i][j] = s[j]
		}
	}

	col, row := [100]bool{}, [100]bool{}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if a[i][j] == '#' {
				col[i] = true
				row[j] = true
			}
		}
	}

	for i := 0; i < h; i++ {
		if col[i] {
			for j := 0; j < w; j++ {
				if row[j] {
					fmt.Print(string(a[i][j]))
				}
			}
			fmt.Println()
		}
	}
}
