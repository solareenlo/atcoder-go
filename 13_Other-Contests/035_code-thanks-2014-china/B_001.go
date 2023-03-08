package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(in, &n)
		r := int(math.Ceil((math.Sqrt(2*float64(n)-1) - 1) / 2))
		a := 2*r*(r+1) + 1
		x := r - (a-n+1)/2
		y := int(math.Pow(-1, float64(a-n+1)) * float64(r-abs(x)))
		fmt.Fprintln(out, x, y)
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
