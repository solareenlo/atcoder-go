package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 200005

var (
	n  int
	a  = make([]int, N)
	sx = make([]int, N)
	sy = make([]int, N)
	mi int
)

func chk() bool {
	t := 0
	for i := 1; i < n; i++ {
		if a[i] <= a[i-1] {
			for ; t != 0 && sx[t] > a[i]; t-- {

			}
			for j := a[i]; j > 0; j-- {
				if sx[t] == j {
					sy[t]++
					if sy[t] < mi {
						break
					}
				} else {
					t++
					sx[t] = j
					sy[t] = 1
					break
				}
				t--
			}
			if sy[t] >= mi || t == 0 {
				return false
			}
		}
	}
	return true
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	ans := 0
	for l, r := 1, n; l <= r; {
		mi = (l + r) >> 1
		if chk() {
			ans = mi
			r = mi - 1
		} else {
			l = mi + 1
		}
	}
	fmt.Println(ans)
}
