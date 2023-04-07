package main

import (
	"bufio"
	"fmt"
	"os"
)

const mod = 1000000007
const inv2 = (mod + 1) / 2

var n int
var bit, fact [200009]int

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n)
	fact[0] = 1
	for i := 1; i <= n; i++ {
		fact[i] = fact[i-1] * i % mod
	}
	for i := 0; i < n; i++ {
		add(i, 1)
	}
	ret := 0
	is := 0
	for i := 0; i < n; i++ {
		var x int
		fmt.Fscan(in, &x)
		x--
		add(x, -1)
		dim := sum(x)
		aug := (is + (n-i-1)*(n-i-2)/2%mod*inv2 + (dim-1)*inv2) % mod
		ret = (ret + aug*dim%mod*fact[n-i-1]) % mod
		is += dim
	}
	fmt.Println((ret + is) % mod)
}

func add(pos, val int) {
	for i := pos + 1; i <= n; i += i & (-i) {
		bit[i] += val
	}
}

func sum(pos int) int {
	ret := 0
	for i := pos; i >= 1; i -= i & (-i) {
		ret += bit[i]
	}
	return ret
}
