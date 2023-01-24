package main

import (
	"bufio"
	"fmt"
	"os"
)

type pair struct {
	x, y int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var T int
	fmt.Fscan(in, &T)
	var b, k, sx, sy, ex, ey int
	s := make([]pair, 4)
	t := make([]pair, 4)
	for T > 0 {
		T--
		fmt.Fscan(in, &b, &k, &sx, &sy, &ex, &ey)
		ans := k * (abs(sx-ex) + abs(sy-ey))
		s[0] = pair{sx / b * b, sy}
		s[1] = pair{sx/b*b + b, sy}
		s[2] = pair{sx, sy / b * b}
		s[3] = pair{sx, sy/b*b + b}
		t[0] = pair{ex / b * b, ey}
		t[1] = pair{ex/b*b + b, ey}
		t[2] = pair{ex, ey / b * b}
		t[3] = pair{ex, ey/b*b + b}
		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				d1 := k * (abs(s[i].x-sx) + abs(s[i].y-sy) + abs(t[j].x-ex) + abs(t[j].y-ey))
				d2 := abs(s[i].x-t[j].x) + abs(s[i].y-t[j].y)
				if s[i].x%b != 0 && s[i].x/b == t[j].x/b && s[i].y/b != t[j].y/b {
					d2 = min(s[i].x%b+t[j].x%b, 2*b-s[i].x%b-t[j].x%b) + abs(s[i].y-t[j].y)
				}
				if s[i].y%b != 0 && s[i].y/b == t[j].y/b && s[i].x/b != t[j].x/b {
					d2 = min(s[i].y%b+t[j].y%b, 2*b-s[i].y%b-t[j].y%b) + abs(s[i].x-t[j].x)
				}
				ans = min(ans, d1+d2)
			}
		}
		fmt.Fprintln(out, ans)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
