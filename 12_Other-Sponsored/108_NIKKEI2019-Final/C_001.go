package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var cntr, cntc [1 << 17]int

	var h, w, k int
	fmt.Fscan(in, &h, &w, &k)
	for i := 0; i < k; i++ {
		var r, c int
		fmt.Fscan(in, &r, &c)
		cntr[r]++
		cntc[c]++
	}
	sumx, sumy := 0, 0
	R, C := 0, 0
	for sumx*2 < h*w-k {
		R++
		sumx += w - cntr[R]
	}
	for sumy*2 < h*w-k {
		C++
		sumy += h - cntc[C]
	}
	ans := 0
	for r := 1; r <= h; r++ {
		ans += (w - cntr[r]) * abs(R-r)
	}
	for c := 1; c <= w; c++ {
		ans += (h - cntc[c]) * abs(C-c)
	}
	fmt.Println(ans)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
