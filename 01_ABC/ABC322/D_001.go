package main

import (
	"fmt"
	"os"
)

var a [3][4][4]int
var b [15][15]int
var c [4][4]int

func p1(t int) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			c[i][j] = a[t][j][3-i]
		}
	}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			a[t][i][j] = c[i][j]
		}
	}
}

func p(t int) {
	if t == 3 {
		for i := 4; i <= 7; i++ {
			for j := 4; j <= 7; j++ {
				if b[i][j] != 1 {
					return
				}
			}
		}
		fmt.Println("Yes")
		os.Exit(0)
	}
	for k := 0; k < 4; k++ {
		for i := 1; i <= 7; i++ {
			for j := 1; j <= 7; j++ {
				for g := 0; g < 4; g++ {
					for h := 0; h < 4; h++ {
						b[i+g][j+h] += a[t][g][h]
					}
				}
				p(t + 1)
				for g := 0; g < 4; g++ {
					for h := 0; h < 4; h++ {
						b[i+g][j+h] -= a[t][g][h]
					}
				}
			}
		}
		p1(t)
	}
}

func main() {
	x := 0
	for g := 0; g < 3; g++ {
		for i := 0; i < 4; i++ {
			var s string
			fmt.Scan(&s)
			for j := 0; j < 4; j++ {
				if s[j] == '#' {
					a[g][i][j] = 1
					x++
				} else {
					a[g][i][j] = 0
				}
			}
		}
	}
	if x == 16 {
		p(0)
	}
	fmt.Println("No")
}
