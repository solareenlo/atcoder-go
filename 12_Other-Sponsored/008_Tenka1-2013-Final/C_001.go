package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var p [4][]int
	var brd [][]int

	var n, m, q int
	fmt.Fscan(in, &n, &m, &q)

	mm := max(n, m)

	for i := 0; i < 4; i++ {
		p[i] = make([]int, n*m+1)
	}
	for i := 1; i <= n*m; i++ {
		for j := 0; j < 4; j++ {
			fmt.Fscan(in, &p[j][i])
		}
	}

	used := make([]bool, n*m+1)
	brd = make([][]int, mm+2)
	for i := range brd {
		brd[i] = make([]int, mm+2)
		for j := range brd[i] {
			brd[i][j] = -1
		}
	}

	var str string

	for i := 0; i < 2*mm-1; i++ {
		for y := 1; y <= mm; y++ {
			x := i - y + 2
			if x < 1 || x > mm {
				continue
			}

			shd := [4]int{15, 15, 15, 15}
			if y == 1 {
				shd[0] = 1
			} else {
				if brd[y-1][x] == -1 {
					continue
				}
				d := brd[y-1][x] & 3
				shd[0] = 1 << (3 - p[(d+2)&3][brd[y-1][x]>>2])
			}

			if x == 1 {
				shd[3] = 1
			} else {
				if brd[y][x-1] == -1 {
					continue
				}
				d := brd[y][x-1] & 3
				shd[3] = 1 << (3 - p[(d+1)&3][brd[y][x-1]>>2])
			}

			if y != n && y != m {
				shd[2] = 14
			}
			if x != n && x != m {
				shd[1] = 14
			}
			if y != 1 && brd[y-1][x+1] == -1 {
				shd[1] = 1
			}
			if x == mm && y == mm {
				shd[1] = 1
				shd[2] = 1
			}

			cand := make([]int, 0)
			for i := 1; i < len(p[0]); i++ {
				if used[i] {
					continue
				}
				for j := 0; j < 4; j++ {
					ok := true
					for k := 0; k < 4; k++ {
						if (shd[k]>>(p[(j+k)&3][i]))&1 == 0 {
							ok = false
						}
					}
					if ok {
						cand = append(cand, i<<2|j)
					}
				}
			}
			if len(cand) == 0 {
				continue
			}

			rand.Shuffle(len(cand), func(i, j int) {
				cand[i], cand[j] = cand[j], cand[i]
			})

			res := -1
			if len(cand) == 1 || (x == 1 && y == 1) {
				res = cand[0]
			} else {
				for i := 0; res == -1 && i < len(cand); i++ {
					if i+1 == len(cand) {
						res = cand[i]
						break
					}

					d := cand[i] & 3
					num := cand[i] >> 2

					if y > 1 {
						adj := brd[y-1][x] >> 2
						adjd := brd[y-1][x] & 3
						fmt.Println("?", adj, (adjd+2)%4+1, num, d+1)
					} else {
						adj := brd[y][x-1] >> 2
						adjd := brd[y][x-1] & 3
						fmt.Println("?", adj, (adjd+1)%4+1, num, (d+3)%4+1)
					}

					fmt.Fscan(in, &str)
					if str[0] == 'y' {
						res = cand[i]
						break
					}
				}
			}

			if res == -1 {
				return
			}
			brd[y][x] = res
			used[res>>2] = true
		}
	}

	if brd[n][m] == -1 {
		tmp := make([][]int, mm+2)
		for i := range tmp {
			tmp[i] = make([]int, mm+2)
			for j := range tmp[i] {
				tmp[i][j] = -1
			}
		}
		for y := 1; y <= n; y++ {
			for x := 1; x <= m; x++ {
				py := x
				px := n + 1 - y
				d := brd[py][px] & 3
				d = (d + 1) & 3
				tmp[y][x] = (brd[py][px] & ^3) | d
			}
		}
		tmp, brd = brd, tmp
	}

	fmt.Println("!")
	for y := 1; y <= n; y++ {
		for x := 1; x <= m; x++ {
			if brd[y][x] == -1 {
				return
			}
			d := brd[y][x] & 3
			d = (4-d)%4 + 1
			if x == m {
				fmt.Println((brd[y][x] >> 2), d)
			} else {
				fmt.Printf("%d %d ", brd[y][x]>>2, d)
			}
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
