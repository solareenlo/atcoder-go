package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N, M int
	fmt.Fscan(in, &N, &M)
	var prime [1 << 19]bool
	for i := 2; i <= N; i++ {
		for j := i * 2; j <= N; j += i {
			prime[j] = true
		}
	}
	var a [1 << 19]int
	var used [1 << 19]bool
	for i := 2; i <= N; i++ {
		if prime[i] == true {
			continue
		}
		X := i
		for X <= N {
			a[X] = int(10000000000.0 * math.Log10(float64(i)))
			used[X] = true
			X *= i
		}
	}
	var c [1 << 19]int
	for i := 0; i < M; i++ {
		fmt.Fscan(in, &c[i])
		if used[c[i]] == true {
			fmt.Fprintln(out, "No")
			return
		}
	}
	fmt.Fprintln(out, "Yes")
	for i := 1; i <= N; i++ {
		if i >= 2 {
			fmt.Fprint(out, " ")
		}
		fmt.Fprint(out, a[i])
	}
	fmt.Fprintln(out)
}
