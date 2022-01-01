package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var h, w int
	fmt.Fscan(in, &h, &w)

	c := make([]string, h)
	for i := 0; i < h; i++ {
		fmt.Fscan(in, &c[i])
	}

	xmin := 1 << 60
	omin := 1 << 60
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if c[i][j] == 'o' {
				break
			}
			if c[i][j] == 'x' {
				xmin = min(xmin, j)
				break
			}
		}
		for j := w - 1; j >= 0; j-- {
			if c[i][j] == 'x' {
				break
			}
			if c[i][j] == 'o' {
				omin = min(omin, w-1-j)
				break
			}
		}
	}

	if xmin != 1<<60 || omin != 1<<60 {
		if xmin >= omin {
			fmt.Println("o")
		} else {
			fmt.Println("x")
		}
		return
	}

	type P struct{ x, y int }
	advance := make([]P, 0)
	orunnum := 0
	xrunnum := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if c[i][j] == '.' {
				continue
			}
			onum := 0
			oright := 0
			for j < w && c[i][j] != 'x' {
				if c[i][j] == 'o' {
					if onum > 0 {
						orunnum += onum * (j - oright - 1)
					}
					onum++
					oright = j
				}
				j++
			}
			xnum := 0
			xl := j
			for j < w && c[i][j] != 'o' {
				if c[i][j] == 'x' {
					if xnum > 0 {
						xrunnum += j - xl - 1 - (xnum - 1)
					}
					xnum++
				}
				j++
			}
			op := oright
			xp := xl
			for xp-op-1 > 2 {
				op++
				xp--
			}
			orunnum += (op - oright) * onum
			xrunnum += (xl - xp) * xnum
			if xp-op-1 == 2 {
				advance = append(advance, P{onum + xnum, xnum})
			}
			j--
		}
	}

	sort.Slice(advance, func(i, j int) bool {
		return advance[i].x > advance[j].x
	})
	for i := 0; i < len(advance); i++ {
		if i%2 == 0 {
			orunnum += advance[i].x - advance[i].y
		} else {
			xrunnum += advance[i].y
		}
	}

	if orunnum > xrunnum {
		fmt.Println("o")
	} else {
		fmt.Println("x")
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
