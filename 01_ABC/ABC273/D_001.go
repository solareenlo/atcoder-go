package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Point struct {
	x, y int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	const N = 200020

	var H, W int
	var C Point
	fmt.Fscan(in, &H, &W, &C.x, &C.y)
	var n int
	fmt.Fscan(in, &n)

	A := make([]Point, N)
	B := make([]Point, N)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &A[i].x, &A[i].y)
		B[i].x = A[i].x
		B[i].y = A[i].y
	}
	tmpA := A[1 : n+1]
	sort.Slice(tmpA, func(i, j int) bool {
		if tmpA[i].x != tmpA[j].x {
			return tmpA[i].x < tmpA[j].x
		}
		return tmpA[i].y < tmpA[j].y
	})
	tmpB := B[1 : n+1]
	sort.Slice(tmpB, func(i, j int) bool {
		if tmpB[i].y != tmpB[j].y {
			return tmpB[i].y < tmpB[j].y
		}
		return tmpB[i].x < tmpB[j].x
	})

	var q int
	fmt.Fscan(in, &q)
	for q > 0 {
		var opt string
		var len int
		fmt.Fscan(in, &opt, &len)
		switch opt {
		case "L":
			i := lowerBoundLR(A[1:n+1], C) + 1
			i--
			if A[i].x != C.x {
				C.y = max(C.y-len, 1)
			} else {
				C.y = max(C.y-len, A[i].y+1)
			}
		case "R":
			i := lowerBoundLR(A[1:n+1], C) + 1
			if A[i].x != C.x {
				C.y = min(C.y+len, W)
			} else {
				C.y = min(C.y+len, A[i].y-1)
			}
		case "U":
			i := lowerBoundUD(B[1:n+1], C) + 1
			i--
			if B[i].y != C.y {
				C.x = max(C.x-len, 1)
			} else {
				C.x = max(C.x-len, B[i].x+1)
			}
		case "D":
			i := lowerBoundUD(B[1:n+1], C) + 1
			if B[i].y != C.y {
				C.x = min(C.x+len, H)
			} else {
				C.x = min(C.x+len, B[i].x-1)
			}
		}
		fmt.Fprintln(out, C.x, C.y)
		q--
	}
}

func lowerBoundLR(a []Point, x Point) int {
	idx := sort.Search(len(a), func(i int) bool {
		if a[i].x != x.x {
			return a[i].x > x.x
		}
		return a[i].y > x.y
	})
	return idx
}

func lowerBoundUD(a []Point, x Point) int {
	idx := sort.Search(len(a), func(i int) bool {
		if a[i].y != x.y {
			return a[i].y > x.y
		}
		return a[i].x > x.x
	})
	return idx
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
