package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 10005

	type pair struct {
		x, y int
	}

	var a [N + 1]int

	var n, r int
	fmt.Fscan(in, &n, &r)
	var fr, to [N][]pair
	for n > 0 {
		n--
		var x1, y1, x2, y2 int
		fmt.Fscan(in, &x1, &y1, &x2, &y2)
		fr[x1] = append(fr[x1], pair{y1, y2})
		to[x2] = append(to[x2], pair{y1, y2})
	}

	s1, s2 := 0, 0
	for i := 0; i < N; i++ {
		for _, y := range fr[i] {
			for j := y.x; j < y.y; j++ {
				a[j]++
				if a[j] == 1 {
					s2++
				}
			}
		}
		for _, y := range to[i] {
			for j := y.x; j < y.y; j++ {
				a[j]--
				if a[j] == 0 {
					s2++
				}
			}
		}
		for j := 0; j < N; j++ {
			if a[j] > 0 {
				s1++
			}
			if j == 0 || a[j-1] == 0 {
				if a[j] > 0 {
					s2++
				}
			}
			if a[j+1] == 0 {
				if a[j] > 0 {
					s2++
				}
			}
		}
	}
	fmt.Println(s1)
	if r == 2 {
		fmt.Println(s2)
	}
}
