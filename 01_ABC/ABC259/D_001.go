package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	n += 2
	const N = 3005
	x := make([]int, N)
	y := make([]int, N)
	fmt.Fscan(in, &x[1], &y[1], &x[n], &y[n])

	r := make([]int, N)
	for i := 2; i < n; i++ {
		fmt.Fscan(in, &x[i], &y[i], &r[i])
	}

	tl := 1
	q := make([]int, N)
	q[tl] = 1
	vst := make([]bool, N)
	vst[1] = true
	hd := 0
	for hd < tl {
		hd++
		k := q[hd]
		for i := 1; i <= n; i++ {
			if !vst[i] {
				u := pow(r[k]-r[i], 2)
				v := pow(x[k]-x[i], 2) + pow(y[k]-y[i], 2)
				w := pow(r[k]+r[i], 2)
				if u <= v && v <= w {
					tl++
					q[tl] = i
					vst[i] = true
				}
			}
		}
	}

	if vst[n] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func pow(a, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = res * a
		}
		a = a * a
		n /= 2
	}
	return res
}
