package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	col := make([]string, n)
	for i := range col {
		fmt.Scan(&col[i])
	}

	row := make([]string, 9)
	for i := 0; i < n; i++ {
		for j := 0; j < 9; j++ {
			row[j] += string(col[i][j])
		}
	}

	cntX, cntO := 0, 0
	for i := 0; i < 9; i++ {
		sep := true
		for j := 0; j < n; j++ {
			if row[i][j] == 'x' {
				cntX++
			}
			if row[i][j] == 'o' {
				if sep {
					cntO++
				}
				sep = false
			} else {
				sep = true
			}
		}
	}

	fmt.Println(cntX + cntO)
}
