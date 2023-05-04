package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var h, w int
	fmt.Fscan(in, &h, &w)
	S := make([]string, h)
	for i := range S {
		fmt.Fscan(in, &S[i])
	}

	row := make([]int, h)
	col := make([]int, w)
	rownum := make([]int, w+1)
	colnum := make([]int, h+1)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if S[i][j] == '#' {
				row[i]++
				col[j]++
			}
		}
	}

	for i := 0; i < h; i++ {
		rownum[row[i]]++
	}
	for j := 0; j < w; j++ {
		colnum[col[j]]++
	}
	r0, r1, c0, c1, ans := 0, 0, 0, 0, 0
	for r0+r1 < h && c0+c1 < w {
		if rownum[c1] == h-r0-r1 {
			ans += rownum[c1] + colnum[r1] - 1
			break
		}
		if rownum[w-c0] == h-r0-r1 {
			ans += rownum[w-c0] + colnum[h-r0] - 1
			break
		}
		if c0+c1 < w && rownum[c1] != 0 {
			r0 += rownum[c1]
			rownum[c1] = 0
			continue
		}
		if c0+c1 < w && rownum[w-c0] != 0 {
			r1 += rownum[w-c0]
			rownum[w-c0] = 0
			continue
		}
		if r0+r1 < h && colnum[r1] != 0 {
			c0 += colnum[r1]
			colnum[r1] = 0
			continue
		}
		if r0+r1 < h && colnum[h-r0] != 0 {
			c1 += colnum[h-r0]
			colnum[h-r0] = 0
			continue
		}
		break
	}
	fmt.Println(r0 + r1 + c0 + c1 + ans)
}
