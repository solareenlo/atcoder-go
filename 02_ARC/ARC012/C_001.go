package main

import "fmt"

var x = [21][21]int{}

func solve() int {
	res := 0
	for i := 1; i < 20; i++ {
		for j := 1; j < 20; j++ {
			for k := 0; k < 9; k++ {
				F1 := (k / 3) - 1
				F2 := (k % 3) - 1
				if k == 4 {
					continue
				}
				cx, cy := i, j
				G := x[i][j]
				d := 5
				if G == 0 {
					continue
				}
				for d > 0 {
					if x[cx][cy] != G {
						break
					}
					cx += F1
					cy += F2
					d--
					if d == 0 {
						res |= G
					}
				}
			}
		}
	}
	return res
}

func main() {
	f, g := 0, 0
	for i := 1; i < 20; i++ {
		var c string
		fmt.Scan(&c)
		c = " " + c
		for j := 1; j < 20; j++ {
			if c[j] == 'o' {
				x[i][j] = 1
				f++
			}
			if c[j] == 'x' {
				x[i][j] = 2
				g++
			}
		}
	}
	if f-g < 0 || f-g > 1 {
		fmt.Println("NO")
		return
	}

	ok := false
	if solve() == 0 {
		ok = true
	}
	for i := 1; i < 20; i++ {
		for j := 1; j < 20; j++ {
			t := x[i][j]
			if f == g && t != 2 {
				continue
			}
			if f != g && t != 1 {
				continue
			}
			x[i][j] = 0
			if solve() == 0 {
				ok = true
			}
			x[i][j] = t
		}
	}
	if ok == false {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}
}
