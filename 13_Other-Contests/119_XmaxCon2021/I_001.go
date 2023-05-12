package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var a, b, c, d int
	fmt.Fscan(in, &a, &b, &c, &d)
	score := 0
	for {
		var e, f int
		fmt.Fscan(in, &e, &f)
		if e == 0 && f == 0 {
			break
		}
		bbminy := min(a, c)
		bbmaxy := max(a, c)
		bbminx := min(b, d)
		bbmaxx := max(b, d)
		bbpy := 0
		if e < bbminy {
			bbpy = bbminy
		} else if e <= bbmaxy {
			bbpy = e
		} else {
			bbpy = bbmaxy
		}
		bbpx := 0
		if f < bbminx {
			bbpx = bbminx
		} else if f <= bbmaxx {
			bbpx = f
		} else {
			bbpx = bbmaxx
		}
		d1 := abs(bbpy-a) + abs(bbpx-b)
		d2 := abs(bbpy-c) + abs(bbpx-d)
		if d1 <= d2 {
			score += abs(a-e) + abs(b-f)
			a = e
			b = f
			if d1 > 0 {
				if c == bbpy {
					if bbpx != d {
						d = d + (bbpx-d)/abs(bbpx-d)*d1
					}
				} else {
					if bbpx != d {
						d = d + (bbpx-d)/abs(bbpx-d)*(d1-1)
					}
				}
				c = bbpy
			}
			fmt.Fprintln(out, 1)
			out.Flush()
		} else {
			score += abs(c-e) + abs(d-f)
			c = e
			d = f
			if d2 > 0 {
				if a == bbpy {
					if bbpx != b {
						b = b + (bbpx-b)/abs(bbpx-b)*d2
					}
				} else {
					if bbpx != b {
						b = b + (bbpx-b)/abs(bbpx-b)*(d2-1)
					}
				}
				a = bbpy
			}
			fmt.Fprintln(out, 2)
			out.Flush()
		}
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
