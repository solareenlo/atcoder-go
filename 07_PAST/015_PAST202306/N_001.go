package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)

	c := make([]int, m)
	for j := 0; j < m; j++ {
		fmt.Fscan(in, &c[j])
	}

	for l := 0; l < 2; l++ {
		t := 0
		p := [2]int{(n - 1) / 2, n / 2}
		q := [2]int{-1, -1}
		o := [2]int{-1, -1}
		for j := 0; j < m; j++ {
			t += c[j]
			for h := 0; h < 2; h++ {
				if q[h] < 0 && t > p[h] {
					q[h] = j
					o[h] = p[h] - (t - c[j])
				}
			}
		}

		s := 0
		for j := 0; j < m; j++ {
			s += j * c[j]
		}
		x := float64(s) / float64(n)
		y := float64(q[0]+q[1]) / float64(2)

		if x < y {
			reverseOrderInt(c)
			continue
		}

		for h := 0; h < 2; h++ {
			s += c[q[h]] - o[h]
			if q[0] == q[1] {
				break
			}
		}
		x = float64(s) / float64(n)
		y += 1

		fmt.Printf("%.12f\n", math.Max(x-y, 0.0))
		break
	}
}

func reverseOrderInt(a []int) {
	n := len(a)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}
