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

	var T int
	fmt.Fscan(in, &T)
	for i := 0; i < T; i++ {
		var xa, ya, ra, xb, yb, rb float64
		fmt.Fscan(in, &xa, &ya, &ra, &xb, &yb, &rb)

		dis := math.Hypot(xa-xb, ya-yb)
		temp := 0.5 / ra
		r1 := temp - 1.0/(dis+ra-rb)
		r2 := temp - 1.0/(dis+ra+rb)

		New := math.Sqrt(r1 * r2)
		r3 := 0.5 / New
		r4 := math.Abs(0.5/(New-r1) - 0.5/(New-r2))

		r5 := (r4 - r3) / 2.0
		angle := math.Asin(r5 / (r3 + r5))
		fmt.Fprintln(out, int(math.Pi/angle))
	}
}
