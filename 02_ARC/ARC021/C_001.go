package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var k, n int
	fmt.Fscan(in, &k, &n)

	A := make([]int, 1<<20)
	D := make([]int, 1<<20)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &A[i], &D[i])
	}

	ok, ng := int(1e14), 0
	for ok-ng > 1 {
		mid := (ok + ng) / 2
		cnt := 0
		for i := 0; i < n; i++ {
			if mid >= A[i] {
				cnt += (mid-A[i])/D[i] + 1
			}
		}
		if cnt >= k {
			ok = mid
		} else {
			ng = mid
		}
	}

	ans, cnt := 0, 0
	for i := 0; i < n; i++ {
		if ok >= A[i] {
			a := (ok - A[i] - 1 + D[i]) / D[i]
			cnt += a
			ans += a*(a-1)/2*D[i] + A[i]*a
		}
	}

	fmt.Println(ans + (k-cnt)*ok)
}
