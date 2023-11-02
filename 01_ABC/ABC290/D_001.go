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

	var t int
	fmt.Fscan(in, &t)
	for t > 0 {
		t--
		var N, D, K int
		fmt.Fscan(in, &N, &D, &K)
		K--
		k := N / gcd(N, D)
		b := K / k
		r := K % k
		fmt.Fprintln(out, (b+r*D%N)%N)
	}
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
