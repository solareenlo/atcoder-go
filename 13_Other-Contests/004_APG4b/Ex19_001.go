package main

import "fmt"

func main() {
	var a [9][9]int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Scan(&a[i][j])
		}
	}

	cntO, cntX := 0, 0
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if a[i][j] == (i+1)*(j+1) {
				cntO++
			} else {
				a[i][j] = (i + 1) * (j + 1)
				cntX++
			}
		}
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Print(a[i][j])
			if j < 8 {
				fmt.Print(" ")
			} else {
				fmt.Println()
			}
		}
	}
	fmt.Println(cntO)
	fmt.Println(cntX)
}
