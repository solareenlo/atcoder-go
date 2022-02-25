package main

import (
	"bufio"
	"fmt"
	"os"
)

const MOD = 998244353

var f = [200002]int{}

func find(x int) int {
	if f[x] == x {
		return f[x]
	}
	f[x] = find(f[x])
	return f[x]
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	for i := 1; i < n+1; i++ {
		f[i] = i
	}

	for i := 1; i < n+1; i++ {
		var x int
		fmt.Fscan(in, &x)
		f[find(x)] = find(i)
	}

	res := 1
	for i := 1; i < n+1; i++ {
		if f[i] == i {
			res = (res << 1) % MOD
		}
	}
	fmt.Println(res - 1)
}
