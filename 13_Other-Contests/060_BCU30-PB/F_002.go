package main

import (
	"bufio"
	"fmt"
	"os"
)

const mod = 1e9 + 7

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, a int
	fmt.Fscan(in, &n)

	pow2 := make([]int, n)
	pow2[0] = 1
	for i := 0; i < n-1; i++ {
		pow2[i+1] = 2 * pow2[i] % mod
	}

	res := 0
	tmp := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a)
		tmp = (tmp * a) % mod
		if i-1 >= 0 {
			tmp += a * pow2[i-1] % mod
		} else {
			tmp += a
		}
		tmp %= mod
		if n-2-i >= 0 {
			res += tmp * pow2[n-2-i] % mod
		} else {
			res += tmp
		}
		res %= mod
	}

	fmt.Println(res)
}
