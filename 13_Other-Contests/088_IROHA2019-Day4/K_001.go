package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	const N = 200005
	const mod = 1000000007

	var n int
	fmt.Fscan(in, &n)
	var p [N + 5]int
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &p[i])
	}
	var inv [N + 5]int
	inv[1] = 1
	ans := 1
	for i := 2; i < n; i++ {
		inv[i] = inv[mod%i] * (mod - mod/i) % mod
		ans = ans * i % mod
	}
	qr := 1
	var q [N + 5]int
	q[qr] = 1
	for i := 2; i <= n; i++ {
		for p[i] > p[q[qr]] {
			qr--
		}
		ans = ans * inv[i-q[qr]] % mod
		qr++
		q[qr] = i
	}
	fmt.Println(ans)
}
