package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, f int
	fmt.Fscan(in, &n, &f)
	var b [5][]int
	for i := 0; i < n; i++ {
		var x int
		fmt.Fscan(in, &x)
		q := x % 5
		if q != 0 {
			b[5-q] = append(b[5-q], x/5+1)
		}
	}
	r := f % 5
	f /= 5
	for i := 1; i < 5; i++ {
		sort.Ints(b[i])
	}
	b3 := make([]int, 100010)
	for i := 0; i < len(b[3]); i++ {
		b3[i+1] = b3[i] + b[3][i]
	}
	m := len(b[3]) + 1
	ans := 0
	s := 0
	for i := 0; i < 4; i++ {
		if len(b[1]) < i {
			continue
		}
		if i > 0 {
			s += b[1][i-1]
		}
		for j := 0; j < 2; j++ {
			if len(b[2]) < j {
				continue
			}
			t := s
			if j == 1 {
				t += b[2][0]
			}
			w := make([]int, len(b[4]))
			copy(w, b[4])
			for k := i; k < len(b[1]); k += 4 {
				if k+3 < len(b[1]) {
					w = append(w, b[1][k]+b[1][k+1]+b[1][k+2]+b[1][k+3])
				}
			}
			for k := j; k < len(b[2]); k += 2 {
				if k+1 < len(b[2]) {
					w = append(w, b[2][k]+b[2][k+1])
				}
			}
			sort.Ints(w)
			if t <= f {
				p := upperBound(b3[:m], f-t) - 1
				if ans < i+j*2+p*3 {
					ans = i + j*2 + p*3
				}
			}
			for k := 0; k < len(w); k++ {
				t += w[k]
				if t > f {
					break
				}
				p := upperBound(b3[:m], f-t) - 1
				if ans < i+j*2+(k+1)*4+p*3 {
					ans = i + j*2 + (k+1)*4 + p*3
				}
			}
		}
	}
	fmt.Println(ans + r)
}

func upperBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] > x
	})
	return idx
}
