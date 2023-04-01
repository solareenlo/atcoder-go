package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var S string
	fmt.Fscan(in, &S)
	X, Y := 0, 0
	var sx, sy int
	if S[0] == 'R' {
		sx = 0
		sy = 1
	} else if S[0] == 'G' {
		sx = 0
		sy = 0
	} else if S[0] == 'B' {
		sx = 0
		sy = 2
	}
	X = sx
	Y = sy
	res := make([]string, 0)
	mnx := sx
	mxx := sx
	mny := sy
	mxy := sy
	dx := [4]int{0, 1, 0, -1}
	dy := [4]int{-1, 0, 1, 0}
	p := [4]string{"GRBG", "BRBR", "BGGR", "BRBR"}
	dir := "LDRU"
	for i := 1; i < len(S); i++ {
		mns := 10000000
		d := 0
		var xx, yy int
		for k := 0; k < 4; k++ {
			xx = X + dx[k]
			yy = Y + dy[k]
			if p[(xx+40000000)%4][(yy+40000000)%4] == S[i] {
				if mns > abs(xx)+abs(yy) {
					mns = abs(xx) + abs(yy)
					d = k
				}
			}
		}
		res = append(res, string(dir[d]))
		X += dx[d]
		Y += dy[d]
		mnx = min(mnx, X)
		mxx = max(mxx, X)
		mny = min(mny, Y)
		mxy = max(mxy, Y)
	}
	fmt.Fprintln(out, mxx-mnx+1, mxy-mny+1)
	for xx := mnx; xx <= mxx; xx++ {
		for yy := mny; yy <= mxy; yy++ {
			fmt.Fprintf(out, "%c", p[(xx+40000000)%4][(yy+40000000)%4])
		}
		fmt.Fprintln(out)
	}
	fmt.Fprintln(out, sx-mnx+1, sy-mny+1)
	fmt.Fprintln(out, strings.Join(res, ""))
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
