package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var f [3][2]float64
	s := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			fmt.Fscan(in, &f[i][j])
			s += int(f[i][j])
		}
	}
	var t [3]int
	x := 0
	var d [2]float64
	for t[0]+t[1]+t[2] != s {
		var e [3]float64
		for i := 0; i < 3; i++ {
			if f[i][0]+f[i][1]-float64(t[i]) != 0.0 {
				e[i] = (100.0*f[i][0] + 50.0*f[i][1]) / (f[i][0] + f[i][1])
			}
		}
		c := -1
		z := -1.0
		for i := 0; i < 3; i++ {
			if z < e[i] {
				c = i
				z = e[i]
			}
		}
		t[c]++
		d[x%2] += z
		x++
	}
	fmt.Println(d[0])
}
