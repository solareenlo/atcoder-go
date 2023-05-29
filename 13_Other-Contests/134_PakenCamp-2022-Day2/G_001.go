package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	type pair struct {
		x, y int
	}

	var N int
	fmt.Fscan(in, &N)
	var query func([]int) bool
	query = func(a []int) bool {
		used := make([]bool, N)
		for _, x := range a {
			used[x] = true
		}
		fmt.Fprintf(out, "? ")
		out.Flush()
		for i := 0; i < N; i++ {
			if used[i] {
				fmt.Fprint(out, 1)
				out.Flush()
			} else {
				fmt.Fprint(out, 0)
				out.Flush()
			}
		}
		fmt.Fprintln(out)
		out.Flush()
		var res string
		fmt.Fscan(in, &res)
		return res == "Yes"
	}
	ans := make([]pair, 0)
	V := make([]int, N)
	for i := range V {
		V[i] = i
	}
	for {
		A := make([]int, 1)
		A[0] = V[0]
		B := make([]int, 0)
		for i := 1; i < len(V); i++ {
			A = append(A, V[i])
			if !query(A) {
				continue
			}
			A = A[:len(A)-1]
			B = append(B, V[i])
		}
		if len(B) == 0 {
			break
		}
		for _, v := range B {
			C := make([]int, len(A))
			copy(C, A)
			for {
				ok := 0
				ng := len(C)
				for ng-ok > 1 {
					mid := (ok + ng) / 2
					tmp := C[mid:]
					D := make([]int, len(tmp))
					copy(D, tmp)
					D = append(D, v)
					if query(D) {
						ok = mid
					} else {
						ng = mid
					}
				}
				ans = append(ans, pair{v, C[ok]})
				C = C[:ok]
				if len(C) == 0 {
					break
				}
				C = append(C, v)
				if !query(C) {
					break
				}
				C = C[:len(C)-1]
			}
		}
		V = B
	}
	fmt.Fprintf(out, "! %d\n", len(ans))
	out.Flush()
	for _, p := range ans {
		fmt.Fprintln(out, p.x+1, p.y+1)
		out.Flush()
	}
}
