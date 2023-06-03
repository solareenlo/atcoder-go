package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var h, w, k int
	fmt.Fscan(in, &h, &w, &k)
	C := make([]string, h)
	for i := 0; i < h; i++ {
		fmt.Fscan(in, &C[i])
	}
	ans := 0
	for i := 0; i < (1 << h); i++ {
		x := i
		y := 0
		z := 0
		for x > 0 {
			if (x & 1) != 0 {
				y++
			}
			x >>= 1
		}
		if y > k {
			continue
		}
		for l := 0; l < h; l++ {
			if ((i >> l) & 1) == 0 {
				for j := 0; j < w; j++ {
					if C[l][j] == '#' {
						z++
					}
				}
			}
		}
		ans = max(ans, z+y*w)
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
