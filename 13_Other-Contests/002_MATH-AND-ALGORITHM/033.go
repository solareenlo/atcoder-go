package main

import (
	"fmt"
	"math"
)

func main() {
	var ax, ay, bx, by, cx, cy float64
	fmt.Scan(&ax, &ay, &bx, &by, &cx, &cy)

	var l float64
	if (bx-ax)*(bx-cx)+(by-ay)*(by-cy) < 0 {
		l = math.Sqrt((bx-ax)*(bx-ax) + (by-ay)*(by-ay))
	} else if (cx-ax)*(cx-bx)+(cy-ay)*(cy-by) < 0 {
		l = math.Sqrt((cx-ax)*(cx-ax) + (cy-ay)*(cy-ay))
	} else {
		l = math.Abs((cx-bx)*(ay-by)-(cy-by)*(ax-bx)) / math.Sqrt((cx-bx)*(cx-bx)+(cy-by)*(cy-by))
	}

	fmt.Println(l)
}
