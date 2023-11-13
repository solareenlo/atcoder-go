package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k, sx, sy, tx, ty int
	fmt.Fscan(in, &n, &k, &sx, &sy, &tx, &ty)
	A := make([]int, 0)
	for i := 0; i < n; i++ {
		var p, q, r, w int
		fmt.Fscan(in, &p, &q, &r, &w)
		if (p*sx+q*sy-r) > 0 && (p*tx+q*ty-r) > 0 {
			A = append(A, 0)
		} else if (p*sx+q*sy-r) < 0 && (p*tx+q*ty-r) < 0 {
			A = append(A, 0)
		} else {
			A = append(A, w)
		}
	}
	sort.Ints(A)
	ans := 0
	for i := 0; i < k; i++ {
		ans += A[i]
	}
	fmt.Println(ans)
}
