package main

import (
	"bufio"
	"fmt"
	"os"
)

var h, w, oddcnt, xcnt, ycnt int
var xmark, ymark [2048]int
var board, quad [2048][2048]int

func touchquad(r, c int) {
	if r < 0 || c < 0 || r >= h-1 || c >= w-1 {
		return
	}
	if quad[r][c] != 0 {
		xmark[c]--
		if xmark[c] == 0 {
			xcnt--
		}
		ymark[r]--
		if ymark[r] == 0 {
			ycnt--
		}
		oddcnt--
	} else {
		if xmark[c] == 0 {
			xcnt++
		}
		xmark[c]++
		if ymark[r] == 0 {
			ycnt++
		}
		ymark[r]++
		oddcnt++
	}
	quad[r][c] ^= 1
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var q int
	fmt.Fscan(in, &h, &w, &q)
	for i := 0; i < h; i++ {
		var buf string
		fmt.Fscan(in, &buf)
		for j := 0; j < w; j++ {
			if buf[j] == '#' {
				board[i][j] = 1
			} else {
				board[i][j] = 0
			}
		}
	}
	for i := 0; i+1 < h; i++ {
		for j := 0; j+1 < w; j++ {
			if board[i][j]^board[i+1][j]^board[i][j+1]^board[i+1][j+1] != 0 {
				touchquad(i, j)
			}
		}
	}
	for i := 0; i < q; i++ {
		if i > 0 {
			var r, c int
			fmt.Fscan(in, &r, &c)
			r--
			c--
			touchquad(r, c)
			touchquad(r-1, c)
			touchquad(r, c-1)
			touchquad(r-1, c-1)
		}
		if oddcnt == 0 || oddcnt == 1 || oddcnt == 2 || oddcnt == 4 {
			if xcnt <= 2 && ycnt <= 2 {
				fmt.Fprintln(out, "Yes")
				continue
			}
		}
		fmt.Fprintln(out, "No")
	}
}
