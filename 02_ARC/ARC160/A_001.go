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

	const N = 100005

	var a, h [N]int

	var n, k int
	fmt.Fscan(in, &n, &k)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
		h[a[i]] = i
	}
	l := 1
	r := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if h[j] < i {
				continue
			}
			if a[i] == j {
				if k-i-(n-i)*(n-i+1)/2 <= 0 {
					break
				}
				k -= i + (n-i)*(n-i+1)/2
			} else {
				k--
				if k == 0 {
					l = i
					r = h[j]
					break
				}
			}
		}
	}
	for i := 1; i < l; i++ {
		fmt.Fprintf(out, "%d ", a[i])
	}
	for i := r; i >= l; i-- {
		fmt.Fprintf(out, "%d ", a[i])
	}
	for i := r + 1; i <= n; i++ {
		fmt.Fprintf(out, "%d ", a[i])
	}
}
