package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var p, q [10]float64

func main() {
	in := bufio.NewReader(os.Stdin)

	for i := 0; i < 6; i++ {
		fmt.Fscan(in, &p[i])
	}
	for i := 0; i < 6; i++ {
		fmt.Fscan(in, &q[i])
	}
	l, r := 0.0, 1.0
	for i := 0; i < 100; i++ {
		ml := (l*2 + r) / 3
		mr := (l + r*2) / 3
		if cal(ml) > cal(mr) {
			l = ml
		} else {
			r = mr
		}
	}
	fmt.Println(cal(l))
}

func cal(mid float64) float64 {
	ans := 0.0
	for i := 0; i < 6; i++ {
		ans += math.Max(p[i]*mid/100, q[i]*(1-mid)/100)
	}
	return ans
}
