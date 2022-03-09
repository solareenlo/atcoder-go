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

	b := [606][606]int{}
	for i := 0; i < h; i++ {
		var s string
		fmt.Fscan(in, &s)
		for j := 0; j < w; j++ {
			b[i+j][i-j+w] = int(s[j] & 1)
		}
	}

	c := [606][606]int{}
	d := [606][606]int{}
	for i := 0; i < h+w; i++ {
		for j := 0; j < h+w; j++ {
			c[i][j+1] = c[i][j] + b[i][j]
			d[i+1][j] = d[i][j] + b[i][j]
		}
	}

	a := 0
	for i := 0; i < h+w; i++ {
		for j := 0; j < h+w; j++ {
			if b[i][j] != 0 {
				for k := 0; k < h+w; k++ {
					if j+k < h+w && b[i][j+k] != 0 {
						tmp1, tmp2 := 0, 0
						if i >= k {
							tmp1 = c[i-k][j+k] - c[i-k][j]
						}
						if i+k < h+w {
							tmp2 = c[i+k][j+k] - c[i+k][j]
						}
						a += tmp1 + tmp2
					}
					if i+k < h+w && b[i+k][j] != 0 {
						tmp1, tmp2 := 0, 0
						if j >= k {
							tmp1 = d[i+k+1][j-k] - d[i][j-k]
						}
						if j+k < h+w {
							tmp2 = d[i+k][j+k] - d[i+1][j+k]
						}
						a += tmp1 + tmp2
					}
				}
			}
		}
	}
	fmt.Println(a)
}
