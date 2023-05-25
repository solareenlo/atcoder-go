package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var a int
	var b string
	fmt.Fscan(in, &a, &b)
	var d [2][2][2][2][2][2]int
	d[0][0][0][0][0][0] = 1
	c := 0
	var e int
	for i := 0; i < a; i++ {
		c = 1 - c
		for j := range d[c] {
			for k := range d[c][j] {
				for l := range d[c][j][k] {
					for m := range d[c][j][k][l] {
						for n := range d[c][j][k][l][m] {
							d[c][j][k][l][m][n] = 0
						}
					}
				}
			}
		}
		if b[i] == 65 {
			e = 0
		} else {
			e = 1
		}
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				for l := 0; l < 2; l++ {
					for m := 0; m < 2; m++ {
						for n := 0; n < 2; n++ {
							f := d[c^1][j][k][l][m][n]
							if f == 0 {
								continue
							}
							if j != 0 {
								g := h(k, e)
								if e != 0 && l != 0 {
									g = 0
								}
								d[c][j^1][g][l][m][n] = (d[c][j^1][g][l][m][n] + f) % 998244353
							} else {
								q := h(k, e)
								r := m
								s := n
								if !(q == 0 && m != 0) && !(q == 1 && n != 0) {
									d[c][0][0][q][0][0] = (d[c][0][0][q][0][0] + f) % 998244353
								}
								if q == 0 {
									r = 1
								} else {
									s = 1
								}
								if q != 0 && n != 0 && e == 0 {
									d[c][0][0][0][0][0] = (d[c][0][0][0][0][0] + f) % 998244353
									r = 1
								}
								d[c][j^1][q][l][r][s] = (d[c][j^1][q][l][r][s] + f) % 998244353
							}
						}
					}
				}
			}
		}
	}
	t := 0
	for u := 0; u < 2; u++ {
		for v := 0; v < 2; v++ {
			for w := 0; w < 2; w++ {
				for x := 0; x < 2; x++ {
					for y := 0; y < 2; y++ {
						z := d[c][u][v][w][x][y]
						if z == 0 {
							continue
						}
						if u == 0 && (v == 0 || y != 0) {
							t = (t + z) % 998244353
						}
					}
				}
			}
		}
	}
	fmt.Println(t)
}

func h(o, p int) int {
	if o == 1 {
		return 0
	}
	return p
}
