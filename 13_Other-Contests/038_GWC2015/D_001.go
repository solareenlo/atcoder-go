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
	v := make([][]int, h)
	for i := range v {
		v[i] = make([]int, w)
	}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			fmt.Printf("1 1 %d %d\n", i+1, j+1)
			fmt.Fscan(in, &v[i][j])
		}
	}

	for i := 0; i < k; i++ {
		var a, b, c, d int
		fmt.Fscan(in, &a, &b, &c, &d)
		a--
		b--
		c--
		d--
		if b > d || (b == d && a < b) {
			a, c = c, a
			b, d = d, b
		}
		if a <= c && b <= d {
			fmt.Println(v[c][d] - v[a][b])
		} else {
			fmt.Println(v[a][b] - v[c][b] - v[c][b] + v[c][d])
		}
	}
}
