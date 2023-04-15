package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, A, B, C, D int
	var s string
	fmt.Fscan(in, &N, &s, &A, &B, &C, &D)
	co, cx := 0, 0
	for i := 0; i < N; i++ {
		if s[i] == 'o' {
			co++
		} else {
			cx++
		}
	}
	if co != A+B+C || cx != A+B+D {
		fmt.Println("No")
		return
	}
	p := 0
	rest := 0
	a := make([]int, 0)
	b := make([]int, 0)
	for i := 1; i <= N; i++ {
		flag := false
		if i == N {
			flag = true
		} else if s[i] == s[i-1] {
			flag = true
		}
		if flag {
			if (i-p)%2 != 0 {
				rest += (i - p) / 2
			} else {
				if s[p] == 'o' {
					a = append(a, (i-p)/2)
				} else {
					b = append(b, (i-p)/2)
				}
			}
			p = i
		}
	}
	sort.Ints(a)
	sort.Ints(b)
	for _, x := range a {
		if A >= x {
			A -= x
		} else {
			if A > 0 {
				x -= A
				A = 0
			}
			rest += x - 1
		}
	}
	for _, y := range b {
		if B >= y {
			B -= y
		} else {
			if B != 0 {
				y -= B
				B = 0
			}
			rest += y - 1
		}
	}
	if A+B <= rest {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
