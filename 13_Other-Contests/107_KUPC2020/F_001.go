package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 100005
	const M = 1000000

	var h, w int
	fmt.Fscan(in, &h, &w)
	var a, b, c, d [N]int
	ans := 0
	for i := 1; i < h; i++ {
		fmt.Fscan(in, &a[i])
		ans += a[i] * w
	}
	for i := 1; i <= w; i++ {
		fmt.Fscan(in, &b[i])
		ans += b[i] * (h - 1)
	}
	for i := 1; i <= h; i++ {
		fmt.Fscan(in, &c[i])
	}
	for i := 1; i < w; i++ {
		fmt.Fscan(in, &d[i])
		ans += d[i]
	}
	var c1, c2 [M*2 + 5]int
	ans += c[1] * (w - 1)
	for i := 1; i < w; i++ {
		t := d[i] - b[i+1]
		c1[t+M]++
		c2[t+M] += t
	}
	for i := 1; i <= 2*M; i++ {
		c1[i] += c1[i-1]
		c2[i] += c2[i-1]
	}
	for i := 2; i <= h; i++ {
		t := c[i] - a[i-1]
		ans += t*c1[M-t] + c2[M-t]
	}
	fmt.Println(ans)
}
